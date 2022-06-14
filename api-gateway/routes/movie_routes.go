package routes

import (
	"net/http"

	"github.com/nilemarezz/my-microservice/api-gateway/handlers"
)

func NewMovieRoute() {

	movieHandler := handlers.NewMovieHandler()

	s := router.PathPrefix("/movies").Subrouter()

	s.HandleFunc("/", movieHandler.GetMovies).Methods(http.MethodGet)
	s.HandleFunc("/", movieHandler.AddMovie).Methods(http.MethodPost)
	s.HandleFunc("/{id:[0-9]+}", movieHandler.GetMovieByID).Methods(http.MethodGet)
}
