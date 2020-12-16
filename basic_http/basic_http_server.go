package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080
	// Call HandleFunc
	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", 8080)
	// Start HTTP Server binding 8080 port
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// Print "Hello World"
	fmt.Fprint(w, "Hello World\n")
}