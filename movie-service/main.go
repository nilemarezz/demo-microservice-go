package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nilemarezz/my-microservice/movie-service/database"
	pb "github.com/nilemarezz/my-microservice/movie-service/proto"
	"github.com/nilemarezz/my-microservice/movie-service/repository"
	"github.com/nilemarezz/my-microservice/movie-service/service"
	"github.com/nilemarezz/my-microservice/movie-service/trace"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
)

func init() {
	initConfig()
	initTimeZone()
	database.ConnectDB()
	trace.InitTracer()
}

func main() {
	// new grpc server with tracing
	s := grpc.NewServer(grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()))

	port := fmt.Sprintf(":%v", 50051)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic("Error listen server")
	}
	log.Println("Server start at port", port)

	repo := repository.NewMovieRepository(database.DB)
	service := service.NewMovieService(repo)
	pb.RegisterMovieServiceServer(s, service)

	err = s.Serve(lis)
	if err != nil {
		panic("Error start grpc server")
	}
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	// if docker set env, then use the env else use in config
	// example  APP_PORT=3000 go run main.go
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		panic("Error load config")
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic("")
	}
	time.Local = ict
}
