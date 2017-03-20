package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/fsgura/testing-go/noauth_gateway/services"
	"github.com/gorilla/mux"
)

var Healthy = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("OK"))
})

/* We will create our catalog of VR experiences and store them in a slice. */
var books = []services.Book {
	services.Book{Id: 1, Title: "Hover Shooters", Slug: "hover-shooters", Description : "Shoot your way to the top on 14 different hoverboards"},
	services.Book{Id: 2, Title: "Ocean Explorer", Slug: "ocean-explorer", Description : "Explore the depths of the sea in this one of a kind underwater experience"},
	services.Book{Id: 3, Title: "Dinosaur Park", Slug : "dinosaur-park", Description : "Go back 65 million years in the past and ride a T-Rex"},
	services.Book{Id: 4, Title: "Cars VR", Slug : "cars-vr", Description: "Get behind the wheel of the fastest cars in the world."},
	services.Book{Id: 5, Title: "Robin Hood", Slug: "robin-hood", Description : "Pick up the bow and arrow and master the art of archery"},
	services.Book{Id: 6, Title: "Real World VR", Slug: "real-world-vr", Description : "Explore the seven wonders of the world in VR"},
}

/* The products handler will be called when the user makes a GET request to the /products endpoint.
   This handler will return a list of products available for users to review */
var BooksHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	// Here we are converting the slice of products to json
	payload, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})

/* The feedback handler will add either positive or negative feedback to the product
   We would normally save this data to the database - but for this demo we'll fake it
   so that as long as the request is successful and we can match a product to our catalog of products
   we'll return an OK status. */
var AddVoteHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	var product services.Book
	vars := mux.Vars(r)
	slug := vars["slug"]
	for _, p := range books {
		if p.Slug == slug {
			product = p
		}
	}
	w.Header().Set("Content-Type", "application/json")
	if product.Slug != "" {
		payload, _ := json.Marshal(product)
		w.Write([]byte(payload))
	} else {
		w.Write([]byte("Book Not Found"))
	}
})

// Here we are implementing the NotImplemented handler. Whenever an API endpoint is hit
// we will simply return the message "Not Implemented"
var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Not Implemented"))
})

