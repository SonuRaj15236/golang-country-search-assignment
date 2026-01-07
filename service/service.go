package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"assignment/model"
)

/*
GetCountry fetches country details from the external
Rest Countries API based on the given country name.

It returns:
- model.Country containing name, capital, currency, and population
- error if API call fails, response is invalid, or data is unavailable
*/
func GetCountry(name string) (model.Country, error) {

	// Initialize empty response struct
	var response model.Country

	// Build API URL dynamically using country name
	api := fmt.Sprintf("https://restcountries.com/v3.1/name/%s?fullText=true", name)

	// Call external API
	resp, er := http.Get(api)
	if er != nil {
		// Return error if API call fails
		return response, fmt.Errorf("can't able to call api - %v", api)
	}

	// Ensure response body is closed after function execution
	defer resp.Body.Close()

	// Validate HTTP status code
	if resp.StatusCode != http.StatusOK {
		return response, fmt.Errorf("data not available for country name: %v ", name)
	}

	// Decode API response into a generic slice
	// Using []any because API response structure is dynamic
	var result []any
	err := json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		// Return error if JSON decoding fails
		return response, fmt.Errorf("can't able to unmarshal body")
	}

	// Extract country name from response
	response.Name = result[0].(map[string]any)["name"].(map[string]any)["common"].(string)

	// Extract capital (first element from capital array)
	response.Capital = result[0].(map[string]any)["capital"].([]any)[0].(string)

	// Extract currency name (dynamic key like INR, USD, etc.)
	currencyMap := result[0].(map[string]any)["currencies"].(map[string]any)
	for _, value := range currencyMap {
		response.Currency = value.(map[string]any)["name"].(string)
	}

	// Extract population (JSON numbers are decoded as float64)
	response.Population = uint64(result[0].(map[string]any)["population"].(float64))

	// Return populated response and nil error
	return response, nil
}
