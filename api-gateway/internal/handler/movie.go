package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	models "github.com/nilemarezz/my-microservice/api-gateway/internal/model"
	utils "github.com/nilemarezz/my-microservice/api-gateway/internal/util"
	pb "github.com/nilemarezz/my-microservice/api-gateway/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc/metadata"
)

type MoveHandler struct {
	client pb.MovieServiceClient
}

func NewMovieHandler() MoveHandler {
	creeds := insecure.NewCredentials()
	// init dial with otel intercepter
	cc, err := grpc.Dial("movie-service:50051", grpc.WithTransportCredentials(creeds),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()))
	if err != nil {
		panic(err)
	}
	return MoveHandler{client: pb.NewMovieServiceClient(cc)}
}

// GetMovies godoc
// @Summary      GetMovies
// @Description  Get All Movies
// @Tags         movies
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Movie
// @Failure      400  {object}  utils.JsonResponse
// @Failure      404  {object}  utils.JsonResponse
// @Failure      500  {object}  utils.JsonResponse
// @Router       /movies/ [get]
func (h MoveHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	// init span - tracing
	tracer := otel.GetTracerProvider().Tracer("http-GetMovies")
	ctx, span := tracer.Start(r.Context(), "http-GetMovies")
	defer span.End()

	fmt.Println("Get Movies")

	// add metadata to ctx
	md := metadata.Pairs(
		"timestamp", time.Now().Format(time.StampNano),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)

	// call grpc service with ctx tracing
	movies, err := h.client.GetAll(ctx, &pb.Empty{})
	if err != nil {
		if grpcError, ok := status.FromError(err); ok {
			utils.ErrorJson(w, grpcError.Err(), http.StatusBadRequest)
			return
		} else {
			utils.ErrorJson(w, err, http.StatusInternalServerError)
			return
		}
	}
	var res []models.Movie
	for _, m := range movies.Movies {

		casts := []models.Cast{}
		for _, c := range m.Cast {
			casts = append(casts, models.Cast{Id: int(c.Id), Name: c.Name, Age: int(c.Age)})
		}
		res = append(res, models.Movie{
			Id:          int(m.Id),
			Name:        m.Name,
			Description: m.Description,
			ScreenDate:  m.ScreenDate,
			Cast:        casts,
		})
	}
	utils.WriteJson(w, http.StatusAccepted, res)
}

// GetMovies godoc
// @Summary      GetMovieByID
// @Description  Get Movie By ID
// @Tags         movies
// @Accept       json
// @Produce      json
// @param id path int true "id of movie to be get"
// @Success      200  {object}   models.Movie
// @Failure      400  {object}  utils.JsonResponse
// @Failure      404  {object}  utils.JsonResponse
// @Failure      500  {object}  utils.JsonResponse
// @Router       /movies/:id [get]
func (h MoveHandler) GetMovieByID(w http.ResponseWriter, r *http.Request) {
	// init span - tracing
	tracer := otel.GetTracerProvider().Tracer("http-GetMovieByID")
	ctx, span := tracer.Start(r.Context(), "http-GetMovieByID")
	defer span.End()
	fmt.Println("GetMovieByID")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// add metadata to ctx
	md := metadata.Pairs(
		"timestamp", time.Now().Format(time.StampNano),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)
	movie, err := h.client.GetByID(ctx, &pb.MovieRequest{Id: uint32(id)})
	if err != nil {
		if grpcError, ok := status.FromError(err); ok {
			utils.ErrorJson(w, grpcError.Err(), http.StatusBadRequest)
			return
		} else {
			utils.ErrorJson(w, err, http.StatusInternalServerError)
			return
		}
	}

	casts := []models.Cast{}
	for _, c := range movie.Movie.Cast {
		casts = append(casts, models.Cast{Id: int(c.Id), Name: c.Name, Age: int(c.Age)})
	}

	res := models.Movie{
		Id:          int(movie.Movie.Id),
		Name:        movie.Movie.Name,
		Description: movie.Movie.Description,
		ScreenDate:  movie.Movie.ScreenDate,
		Cast:        casts,
	}
	utils.WriteJson(w, http.StatusAccepted, res)

}

func (h MoveHandler) AddMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Movie Add")
}
