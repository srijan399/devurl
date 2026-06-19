package main

import (
	"fmt"
	"log"
	"net/http"
	"urlshort/controllers"
	"urlshort/internal/db"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect postgres
	db.ConnectMain()

	// Homepage
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hey, welcome to URLShortener!")
	})

	// Shorten POST Endpoint
	http.HandleFunc("/shorten", controllers.ShortenURL)

	// Fetch ORIGINAL URL
	http.HandleFunc("/{code}", controllers.Fetch)

	fmt.Println("Server running on :8090")
	http.ListenAndServe(":8090", nil)
}
