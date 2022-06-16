package route

import (
	"net/http"

	handlers "github.com/nilemarezz/my-microservice/api-gateway/internal/handler"
	pb "github.com/nilemarezz/my-microservice/api-gateway/proto"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAuthRoute() {

	creeds := insecure.NewCredentials()
	// init dial with otel intercepter
	cc, err := grpc.Dial("auth-service:50052", grpc.WithTransportCredentials(creeds),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()))
	if err != nil {
		panic(err)
	}

	authHandler := handlers.NewAuthHandler(pb.NewAuthServiceClient(cc))

	s := router.PathPrefix("/auth").Subrouter()

	s.HandleFunc("/login", authHandler.Login).Methods(http.MethodPost)
	s.HandleFunc("/signup", authHandler.Signup).Methods(http.MethodPost)

}
