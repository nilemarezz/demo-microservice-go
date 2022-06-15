package route

import (
	"net/http"

	handlers "github.com/nilemarezz/my-microservice/api-gateway/internal/handler"
)

func NewMovieRoute() {

	movieHandler := handlers.NewMovieHandler()

	s := router.PathPrefix("/movies").Subrouter()

	s.HandleFunc("/", movieHandler.GetMovies).Methods(http.MethodGet)
	s.HandleFunc("/", movieHandler.AddMovie).Methods(http.MethodPost)
	s.HandleFunc("/{id:[0-9]+}", movieHandler.GetMovieByID).Methods(http.MethodGet)
}
