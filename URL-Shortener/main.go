package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Link stores one URL
type Link struct {
	ID      int    `json:"id"`
	LongURL string `json:"long_url"`
}

// This slice works as our temporary database
var links []Link

// ID for the next URL
var nextID = 1

func shortenURL(w http.ResponseWriter, r *http.Request) {

	// Store the JSON data sent by the user
	var data Link

	// Read JSON from request
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Check if URL was provided
	if data.LongURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	// Give the URL a unique ID
	data.ID = nextID

	// Increase ID for the next URL
	nextID++

	// Store the URL
	links = append(links, data)

	// Create short URL
	shortURL := "http://localhost:8080/" + strconv.Itoa(data.ID)

	// Prepare response
	response := map[string]string{
		"short_url": shortURL,
		"long_url":  data.LongURL,
	}

	// Tell browser/client that response is JSON
	w.Header().Set("Content-Type", "application/json")

	// Send response
	json.NewEncoder(w).Encode(response)
}

func redirectURL(w http.ResponseWriter, r *http.Request) {

	// Get ID from URL
	idText := r.PathValue("id")

	// Convert ID from string to integer
	id, err := strconv.Atoi(idText)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Search for the URL
	for i := 0; i < len(links); i++ {

		if links[i].ID == id {

			// Redirect to original URL
			http.Redirect(
				w,
				r,
				links[i].LongURL,
				http.StatusFound,
			)

			return
		}
	}

	// ID was not found
	http.Error(w, "URL not found", http.StatusNotFound)
}

func main() {

	// Create short URL
	http.HandleFunc("POST /shorten", shortenURL)

	// Redirect short URL
	http.HandleFunc("GET /{id}", redirectURL)

	// Home page
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "URL Shortener")
	})

	fmt.Println("Server running on http://localhost:8080")

	// Start server
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}
