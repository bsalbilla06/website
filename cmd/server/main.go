package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bsalbilla06/website/internal/handlers"
)

func main() {
	// Set up handlers
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.HandleFunc("/experience", handlers.ExperienceHandler)
	http.HandleFunc("/contact", handlers.ContactHandler)

	// Serve static files
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
