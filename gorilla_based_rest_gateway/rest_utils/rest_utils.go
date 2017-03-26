package rest_utils

import (
	"net/http"
)

type HttpError struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      int    `json:"-"`
}

func NewHttpError(description string, status int) *HttpError {
	return &HttpError{
		Title:       http.StatusText(status),
		Description: description,
		Status:      status,
	}
}
