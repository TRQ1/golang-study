package main

import (
		"encoding/json"
		"fmt"
		"log"
		"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

type helloWorldHandler struct {}

type validationHander struct {
	next http.Handler
}

func main() {

		port := 8080

		handler := newValidationHandler(newHelloWorldHandler())

		http.Handle("/helloworld", handler)


		log.Printf("Server starting on port %v\n", port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}

func newValidationHandler(next http.Handler) http.Handler {
		return validationHander{next: next}
}

func (h validationHander) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
		var request helloWorldRequest
		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&request)
		if err != nil {
				http.Error(rw, "Bad request", http.StatusBadRequest)
				return
		}

		h.next.ServeHTTP(rw, r)
}
func newHelloWorldHandler() http.Handler {
		return helloWorldHandler{}
}

func (h helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
		response := helloWorldResponse{Message: "Hello"}

		encoder := json.NewEncoder(rw)
		encoder.Encode(response)
}
