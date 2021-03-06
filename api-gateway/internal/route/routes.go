package route

import (
	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
)

var router *mux.Router

func NewRouter() *mux.Router {
	router = mux.NewRouter()
	NewMovieRoute()
	NewAuthRoute()
	router.Use(otelmux.Middleware("api-gateway"))
	return router
}
