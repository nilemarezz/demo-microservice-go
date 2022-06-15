package service

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/nilemarezz/my-microservice/movie-service/internal/model"
	"github.com/nilemarezz/my-microservice/movie-service/internal/repository"
	pb "github.com/nilemarezz/my-microservice/movie-service/proto"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	otelCodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type movieService struct {
	pb.MovieServiceServer
	repo repository.MovieRepository
}

func NewMovieService(repo repository.MovieRepository) pb.MovieServiceServer {
	return movieService{repo: repo}
}

func (s movieService) GetAll(ctx context.Context, req *pb.Empty) (*pb.MoviesResponse, error) {
	// start trace
	tracer := otel.GetTracerProvider().Tracer("movie-service")
	ctx, span := tracer.Start(ctx, "service/GetAll")
	defer span.End()

	movies, err := s.repo.All(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var res []*pb.Movie
	for _, m := range movies {
		cast := getCast(m.Cast, ctx, m.Id)
		res = append(res, &pb.Movie{
			Id:          uint32(m.Id),
			Name:        m.Name,
			Description: m.Description,
			ScreenDate:  m.ScreenDate,
			Cast:        cast,
		})
	}
	return &pb.MoviesResponse{Movies: res}, nil
}

func (s movieService) GetByID(ctx context.Context, req *pb.MovieRequest) (*pb.MovieResponse, error) {
	tracer := otel.GetTracerProvider().Tracer("movie-service")
	_, span := tracer.Start(ctx, "service/GetByID")
	defer span.End()

	movie, err := s.repo.ByID(req.Id, ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			span.RecordError(err)
			span.SetStatus(otelCodes.Error, err.Error())
			return nil, status.Error(codes.NotFound, "Item Not Found")
		}
		return nil, err
	}
	var cast []*pb.Cast
	for _, c := range movie.Cast {
		cast = append(cast, &pb.Cast{Id: uint32(c.Id), Name: c.Name, Age: uint32(c.Age)})
	}
	res := &pb.Movie{
		Id:          uint32(movie.Id),
		Name:        movie.Name,
		Description: movie.Description,
		ScreenDate:  movie.ScreenDate,
		Cast:        cast,
	}
	return &pb.MovieResponse{Movie: res}, nil
}

func (s movieService) AddMovie(ctx context.Context, req *pb.Movie) (*pb.AddMovieResponse, error) {
	res := &pb.AddMovieResponse{
		Id:      112,
		Success: true,
	}

	return res, nil
}

func getCast(casts []model.Cast, ctx context.Context, id int) []*pb.Cast {

	tracer := otel.GetTracerProvider().Tracer("movie-service")
	_, span := tracer.Start(ctx, "service/getCast")
	span.SetAttributes(attribute.String("actorId", strconv.Itoa(id)))
	defer span.End()

	var cast []*pb.Cast
	for _, c := range casts {
		cast = append(cast, &pb.Cast{Id: uint32(c.Id), Name: c.Name, Age: uint32(c.Age)})
	}
	time.Sleep(500 * time.Millisecond)
	return cast
}
