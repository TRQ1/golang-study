package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

type helloWorldResponse struct {
	Message string `json:"message"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

const port = 8080

func main() {
		go server()
}

func server() {
		http.HandleFunc("/helloworld", helloWorldHandler)

		log.Printf("Server starting on port %v\n", port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
		var request helloWorldRequest
		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&request)
		if err != nil {
				http.Error(w, "Bad request", htpp.StatusBadRequest)
		}

		response := helloWorldResponse{Message: "Hello" + request.Name}

		encoder := json.NewEncoder(w)
		encoder.Encode(response)

}