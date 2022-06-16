package main

import (
	"log"
	"net/http"
	"time"

	"github.com/nilemarezz/my-microservice/api-gateway/internal/route"
	"github.com/nilemarezz/my-microservice/api-gateway/tools/trace"
	"github.com/rs/cors"

	// tracing

	_ "github.com/nilemarezz/my-microservice/api-gateway/docs" // docs is generated by Swag CLI, you have to import it.
	httpSwagger "github.com/swaggo/http-swagger"               // http-swagger middleware
)

var port = ":5000"

func init() {
	trace.InitTracer()
	initTimeZone()
}

// @title           Demo-microservice
// @version         1.0
// @description     This is a my demo microservice for experimental.

// @contact.name   Matas Paosriwong
// @contact.email  nilenon@gmail.com
func main() {
	router := route.NewRouter()
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	// cors
	// ch := gohandler.CORS(gohandler.AllowedOrigins([]string{"*"}))

	srv := &http.Server{
		Addr:         port,
		Handler:      c.Handler(router),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Printf("Application at port %v", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic("")
	}
	time.Local = ict
}
