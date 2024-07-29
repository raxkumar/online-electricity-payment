package service

import (
	"encoding/json"
	"net/http"
)

type Test struct {
	Name        string `json:"name"`
	Subject     string `json:"subject,omitempty"`
	Description string `json:"description,omitempty"`
}

// HandleTestRoute is the handler function for the "/test" route
func HandleTestRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, this is the response for /test"))
}

// TestMethod is the handler function for a route that accepts a Test struct in JSON format
func TestMethod(w http.ResponseWriter, r *http.Request) {
	var test Test

	// Decode the JSON request body into the Test struct
	err := json.NewDecoder(r.Body).Decode(&test)
	if err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	// You can now use the 'test' variable containing the decoded JSON data
	// For example, printing it or returning it in the response
	response := map[string]interface{}{
		"message": "Received Test data",
		"test":    test,
	}

	// Encode the response into JSON and write it to the response writer
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}
}
