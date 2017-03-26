package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"github.com/fsgura/testing-go/gorilla_based_rest_gateway/rest_utils"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var crowd []Person

func GetPersonById(ctx context.Context, w http.ResponseWriter, req *http.Request) *rest_utils.HttpError {
	params := mux.Vars(req)
	for _, item := range crowd {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func GetCrowd(ctx context.Context, w http.ResponseWriter, req *http.Request) *rest_utils.HttpError {
	json.NewEncoder(w).Encode(crowd)
}

func CreatePerson(ctx context.Context, w http.ResponseWriter, req *http.Request) *rest_utils.HttpError {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	crowd = append(crowd, person)
	json.NewEncoder(w).Encode(crowd)
}

func DeletePerson(ctx context.Context, w http.ResponseWriter, req *http.Request) *rest_utils.HttpError {
	params := mux.Vars(req)
	for index, item := range crowd {
		if item.ID == params["id"] {
			crowd = append(crowd[:index], crowd[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(crowd)
}

