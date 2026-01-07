package handler

import (
	"fmt"
	"net/http"
	"sync"

	"assignment/service"
)

// cache is an in-memory concurrent-safe map
// used to store country responses to avoid
// repeated external API calls for the same country
var cache sync.Map

/*
Index is an HTTP handler function that:
1. Reads the country name from query parameters
2. Checks if the country data exists in in-memory cache
3. Returns cached data if available
4. Calls service layer to fetch data if not cached
5. Stores the response in cache and returns it to the client
*/
func Get(w http.ResponseWriter, r *http.Request) {

	// Read 'name' query parameter from request URL
	// Example: /country?name=India
	name := r.URL.Query().Get("name")

	// Check if country data already exists in cache
	val, ok := cache.Load(name)
	if ok {
		// Cache hit: return cached response
		fmt.Println("Exists")
		fmt.Fprintf(w, "%v", val)
		return
	}

	// Cache miss: data not found in cache
	fmt.Println("Not Exists")

	// Call service layer to fetch country details
	response, err := service.GetCountry(name)
	if err != nil {
		// Return error response if service call fails
		fmt.Fprintf(w, "%v", err)
		return
	}

	// Store fetched response in cache for future requests
	cache.Store(response.Name, response)

	// Write final response back to client
	fmt.Fprintf(w, "%v", response)
}
