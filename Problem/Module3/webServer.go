package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	/* 3b
	Output:
	Server listening on port 8080..
	*/
	// Define the route handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!") // Send "Hello World!" as the response
	}

	// Create a new_server mux
	mux := http.NewServeMux()

	// Register the route handler function for the "/" route
	mux.HandleFunc("/", handler)

	// Create a new HTTP server
	server := &http.Server{
		Addr:    ":8080", // Listen to localhost:8080
		Handler: mux,
	}
	// Start the server
	log.Println("Server listening on port 8080..")
	log.Fatal(server.ListenAndServe())
}
