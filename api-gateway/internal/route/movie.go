package route

import (
	"net/http"

	handlers "github.com/nilemarezz/my-microservice/api-gateway/internal/handler"
	"github.com/nilemarezz/my-microservice/api-gateway/internal/middleware"
	pb "github.com/nilemarezz/my-microservice/api-gateway/proto"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewMovieRoute() {

	creeds := insecure.NewCredentials()
	// init dial with otel intercepter
	ccMovie, err := grpc.Dial("movie-service:50051", grpc.WithTransportCredentials(creeds),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()))
	if err != nil {
		panic(err)
	}

	ccAuth, err := grpc.Dial("auth-service:50052", grpc.WithTransportCredentials(creeds),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()))
	if err != nil {
		panic(err)
	}

	movieHandler := handlers.NewMovieHandler(pb.NewMovieServiceClient(ccMovie))

	s := router.PathPrefix("/movies").Subrouter()

	s.HandleFunc("/", movieHandler.GetMovies).Methods(http.MethodGet)
	s.HandleFunc("/", movieHandler.AddMovie).Methods(http.MethodPost)
	s.HandleFunc("/{id:[0-9]+}", movieHandler.GetMovieByID).Methods(http.MethodGet)

	authMiddleware := middleware.NewAuthMiddleware(pb.NewAuthServiceClient(ccAuth))
	s.Use(authMiddleware.VerifyToken)
}
