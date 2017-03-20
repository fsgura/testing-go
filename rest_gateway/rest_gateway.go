package main

import (
	"github.com/fsgura/testing-go/rest_gateway/routers"
	"github.com/fsgura/testing-go/rest_gateway/settings"
	"github.com/codegangsta/negroni"
	"net/http"
	"fmt"
)

func main() {
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	fmt.Println("Going to start an api gateway on local ip port 5000 ...")
	http.ListenAndServe(":5000", n)
}