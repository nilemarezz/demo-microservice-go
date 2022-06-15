package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/nilemarezz/my-microservice/auth-service/connection"
	"github.com/nilemarezz/my-microservice/auth-service/internal/repository"
	"github.com/nilemarezz/my-microservice/auth-service/internal/service"
	pb "github.com/nilemarezz/my-microservice/auth-service/proto"
	"github.com/nilemarezz/my-microservice/auth-service/trace"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func init() {
	initTimeZone()
	connection.ConnectMySQLDB()
	trace.InitTracer()
}

func main() {

	// new grpc server with tracing
	s := grpc.NewServer(grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()))

	port := fmt.Sprintf(":%v", 50052)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic("Error listen server")
	}

	repo := repository.NewMovieRepository(connection.MySqlDB)
	pb.RegisterAuthServiceServer(s, service.NewUserService(repo))

	log.Println("Server start at port", port)

	err = s.Serve(lis)
	if err != nil {
		panic("Error start grpc server")
	}

}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic("")
	}
	time.Local = ict
}
