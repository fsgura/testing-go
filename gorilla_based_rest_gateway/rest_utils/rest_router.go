package rest_utils

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
	"github.com/gorilla/mux"
)

var crowd []Person

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) *HttpError

func GetPersonById(w http.ResponseWriter, req *http.Request) *HttpError {
	params := mux.Vars(req)
	for _, item := range crowd {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func GetCrowd(w http.ResponseWriter) *HttpError {
	json.NewEncoder(w).Encode(crowd)
}

func CreatePerson(w http.ResponseWriter, req *http.Request) *HttpError {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	crowd = append(crowd, person)
	json.NewEncoder(w).Encode(crowd)
}

func DeletePerson(w http.ResponseWriter, req *http.Request) *HttpError {
	params := mux.Vars(req)
	for index, item := range crowd {
		if item.ID == params["id"] {
			crowd = append(crowd[:index], crowd[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(crowd)
}

func NewRouter(ctx context.Context, routes map[string]map[string]Handler ) *mux.Router {
	r := mux.NewRouter()

	for method, mappings := range routes {
		for route, fct := range mappings {
			localFct := fct
			wrap := func(w http.ResponseWriter, r *http.Request) {
				log.WithFields(log.Fields{"method": r.Method, "uri": r.RequestURI}).Info("HTTP request received")

				err := localFct(ctx, w, r)
				if err != nil {
					log.WithFields(log.Fields{"method": r.Method, "uri": r.RequestURI}).Info(err.Description)
					w.Header().Set("Content-Type", "application/json; charset=utf-8")
					w.Header().Set("X-Content-Type-Options", "nosniff")
					w.WriteHeader(err.Status)
					enc := json.NewEncoder(w)
					enc.Encode(err)
					return
				}
			}

			r.Path(route).Methods(method).HandlerFunc(wrap)
		}
	}
	return r
}

