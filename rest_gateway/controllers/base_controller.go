package controllers

import (
	"net/http"
)

func BaseRouterController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Write([]byte("Welcome to the rest_gateway base API"))
}