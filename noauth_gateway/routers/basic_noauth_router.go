package routers

import (
	"github.com/fsgura/testing-go/noauth_gateway/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func SetNoAuthRoutes(router *mux.Router) *mux.Router {

	// On the default page we will simply serve our static index page.
	router.Handle("/", http.FileServer(http.Dir("./noauth_gateway/html/content_pages/")))

	// We will setup our server so we can serve static assets like images, css from the /static/{file} route
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./noauth_gateway/html/content_static/"))))


	router.Handle("/status",controllers.Healthy).Methods("GET")

	router.Handle("/books", controllers.BooksHandler).Methods("GET")

	router.Handle("/book/{slug}/vote", controllers.AddVoteHandler).Methods("POST")

	return router
}
