package models

import (
		"net/http"
)

type validationContextKey string

type helloWorldResponse struct {
	Message string `json:"message"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

type validationHandler struct {
	next http.Handler
}