package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"urlshort/internal/db"
	"urlshort/types"
	"urlshort/utils"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func ShortenURL(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)

	if err != nil {
		log.Fatal("Error reading request body.")
	}

	var originalURL types.URL
	err = json.Unmarshal(body, &originalURL)

	if !utils.Valid(originalURL.Url) {
		http.Error(res, "URL is not valid", http.StatusBadRequest)
		return
	}

	var shortCode string
	err = db.DB.QueryRow(
		context.Background(),
		`SELECT short_code FROM urls WHERE original_url = $1`,
		originalURL.Url,
	).Scan(&shortCode)

	if err == pgx.ErrNoRows {
		shortCode = utils.GenerateShortCode()
		_, err = db.DB.Exec(
			context.Background(),
			`INSERT INTO urls (short_code, original_url) VALUES ($1, $2)`,
			shortCode, originalURL.Url,
		)
		if err != nil {
			http.Error(res, "failed to insert", http.StatusInternalServerError)
			return
		}
	} else if err != nil {
		http.Error(res, "db error", http.StatusInternalServerError)
		return
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	gateway := os.Getenv("GATEWAY_URL")
	targetURL := fmt.Sprintf(
		"%v%v", gateway, shortCode,
	)

	fmt.Fprintf(res, "Short URL: %s", targetURL)
}
