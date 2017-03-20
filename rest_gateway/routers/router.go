package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetBasicNoAuthRoutes(router)
	router = SetAuthRoutes(router)
	return router
}