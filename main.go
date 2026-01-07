package main

import (
	"fmt"
	"net/http"

	"assignment/handler"
)

/*
main is the entry point of the application.
It performs the following steps:
1. Creates a new HTTP router using Gorilla Mux
2. Configures CORS policies
3. Registers application routes
4. Starts the HTTP server on port 8000
*/
func main() {

	http.HandleFunc("/api/countries/search", handler.Get)

	// Log server start message
	fmt.Println("Server Starting at - http://localhost:8000")

	// Start HTTP server on port 8000 with CORS-enabled router
	// handlers.CORS wraps the router with CORS middleware
	if err := http.ListenAndServe(":8000", nil); err != nil {
		// Log error if server fails to start
		fmt.Println(err.Error())
	}
}
