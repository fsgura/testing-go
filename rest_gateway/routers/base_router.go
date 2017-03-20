package routers

import (
	"github.com/fsgura/testing-go/rest_gateway/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetBasicNoAuthRoutes(router *mux.Router) *mux.Router {
	router.Handle("/test/hello",
		negroni.New(
			negroni.HandlerFunc(controllers.BaseRouterController),
		)).Methods("GET")

	return router
}