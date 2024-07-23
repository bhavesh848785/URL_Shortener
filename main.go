package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

var urlDB = make(map[string]URL)

func generateShortURL(originalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(originalURL))
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)
	return hash[:8]
}

func createURL(originalURL string) string {
	shortURL := generateShortURL(originalURL)
	id := shortURL
	urlDB[id] = URL{
		ID:           id,
		OriginalURL:  originalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}

func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}

func RootPageURL(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortURL := createURL(data.URL)
	response := struct {
		ShortURL string `json:"short_url"`
	}{ShortURL: shortURL}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/redirect/")
	fmt.Println("Redirect request for ID:", id) // Debug statement

	url, err := getURL(id)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusNotFound)
		fmt.Println("Error retrieving URL:", err) // Debug statement
		return
	}
	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func main() {
	fmt.Println("Starting URL shortener...")

	// Register the handler function to handle all requests to the root URL ("/")
	http.HandleFunc("/", RootPageURL)
	http.HandleFunc("/shorten", ShortURLHandler)
	http.HandleFunc("/redirect/", redirectURLHandler)

	// Start the HTTP server on port 3000
	fmt.Println("Starting server on port 3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error on starting server:", err)
	}
}
