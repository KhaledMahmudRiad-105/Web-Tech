package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Link stores information about one URL
type Link struct {
	ID      int    `json:"id"`
	LongURL string `json:"long_url"`
}

// This slice temporarily stores our URLs
var links []Link

// This gives every new URL a new ID
var nextID = 1

// This function creates a short URL
func shortenURL(w http.ResponseWriter, r *http.Request) {

	// Read the URL sent by the user
	var data Link

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Give the URL an ID
	data.ID = nextID

	// Increase ID for the next URL
	nextID++

	// Store the URL
	links = append(links, data)

	// Create the short URL
	shortURL := "http://localhost:8080/" + strconv.Itoa(data.ID)

	// Send the result back
	response := map[string]string{
		"short_url": shortURL,
		"long_url":  data.LongURL,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func main() {

	// POST /shorten creates a short URL
	http.HandleFunc("POST /shorten", shortenURL)

	// Home page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "URL Shortener")
	})

	fmt.Println("Server running on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
