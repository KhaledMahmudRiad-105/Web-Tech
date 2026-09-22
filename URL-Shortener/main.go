package main

import (
	"fmt"
	"net/http"
)

// Link stores information about one URL
type Link struct {
	ID      int
	LongURL string
}

// This slice will temporarily store our URLs
var links []Link

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "URL Shortener")
	})
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
