package main

import (
		"context"
		"encoding/json"
		"fmt"
		"log"
		"net/http"
		"time"
		"testing"
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

type helloWorldHandler struct {
}

func server() {
		port := 8080

		handler := newValidationHandler(newHelloWorldHandler())
		http.Handle("/helloworld", handler)

		log.Printf("Server starting on port %v\n", port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))	
}

func newValidationHandler(next http.Handler) http.Handler {
		return validationHandler{next: next}
}

func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
		var request helloWorldRequest
		decoder := json.NewDecoder(r.Body)
		
		err := decoder.Decode(&request)
		if err != nil {
				http.Error(rw, "Bad Request", http.StatusBadRequest)
				return
		}

		c := context.WithValue(r.Context(), validationContextKey("name"), request.Name)
		r = r.WithContext(c)

		h.next.ServeHTTP(rw, r)
}

func newHelloWorldHandler() http.Handler {
		return helloWorldHandler{}
}

func (h helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	name := r.Context().Value(validationContextKey("name")).(string)
	response := helloWorldResponse{Message: "Hello " + name}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

func fetchGoogle(t *testing.T) {
	r, _ := http.NewRequest("GET", "https://google.com", nil)

	timeoutRequest, cancelFunc := context.WithTimeout(r.Context(), 1*time.Millisecond)
	defer cancelFunc()

	r = r.WithContext(timeoutRequest)

	_, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("Error:", err)
	}
}