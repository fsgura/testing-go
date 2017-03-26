
package main


import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Product struct {
	Id          int
	Name        string
	Slug        string
	Description string
}

var products = []Product{
	{Id: 1, Name: "Hover Shooters", Slug: "hover-shooters", Description: "Shoot your way to the top on 14 different hoverboards"},
	{Id: 2, Name: "Ocean Explorer", Slug: "ocean-explorer", Description: "Explore the depths of the sea in this one of a kind underwater experience"},
	{Id: 3, Name: "Dinosaur Park", Slug: "dinosaur-park", Description: "Go back 65 million years in the past and ride a T-Rex"},
	{Id: 4, Name: "Cars VR", Slug: "cars-vr", Description: "Get behind the wheel of the fastest cars in the world."},
	{Id: 5, Name: "Robin Hood", Slug: "robin-hood", Description: "Pick up the bow and arrow and master the art of archery"},
	{Id: 6, Name: "Real World VR", Slug: "real-world-vr", Description: "Explore the seven wonders of the world in VR"},
}

func main() {

	err := godotenv.Load("auth0.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("views/")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	// Auth0 Token
	r.Handle("/status", StatusHandler).Methods("GET")
	r.Handle("/products", jwtMiddleware.Handler(ProductsHandler)).Methods("GET")
	r.Handle("/products/{slug}/feedback", jwtMiddleware.Handler(AddFeedbackHandler)).Methods("POST")

	// Manual Token
	r.Handle("/get-token", GetTokenHandler).Methods("GET")
	r.Handle("/manual/products/", ValidateToken.Handler(ProductsHandler)).Methods("GET")
	r.Handle("/manual/products/{slug}/feedback", ValidateToken.Handler(AddFeedbackHandler)).Methods("POST")

	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}

var mySigningKey = []byte("secret")

// Handlers
var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["name"] = "Fabrizius"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(tokenString))
})

var ValidateToken = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		decoded := []byte(os.Getenv("AUTH0_CLIENT_SECRET"))
		if len(decoded) == 0 {
			return nil, errors.New("Missing Client Secret")
		}
		return decoded, nil
	},
})

//var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("Not Implemented"))
//})

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
})

var ProductsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	payload, _ := json.Marshal(products)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})

var AddFeedbackHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var product Product
	vars := mux.Vars(r)
	slug := vars["slug"]

	for _, p := range products {
		if p.Slug == slug {
			product = p
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if product.Slug != "" {
		payload, _ := json.Marshal(product)
		w.Write([]byte(payload))
	} else {
		w.Write([]byte("Product Not Found"))
	}
})