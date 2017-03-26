package main

import (
	"github.com/gorilla/handlers"
	"github.com/fsgura/testing-go/noauth_gateway/routers"
	"fmt"
	"net/http"
	"os"
)

func main() {
	router := routers.InitRoutes()
	fmt.Println("Going to start an api gateway on local ip port 5000 ...")
	fmt.Println("The /status endpoint will answer with the healthcheck of this RESTful app")
	http.ListenAndServe(":5000", handlers.LoggingHandler(os.Stdout, router))
}
