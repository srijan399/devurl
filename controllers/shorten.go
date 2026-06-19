package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"urlshort/internal/db"
	"urlshort/types"
	"urlshort/utils"

	"github.com/jackc/pgx/v5"
)

func ShortenURL(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "error reading request body", http.StatusBadRequest)
		return
	}

	var originalURL types.URL
	if err := json.Unmarshal(body, &originalURL); err != nil {
		http.Error(res, "invalid JSON body", http.StatusBadRequest)
		return
	}

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

	gateway := os.Getenv("GATEWAY_URL")
	targetURL := fmt.Sprintf("%v%v", gateway, shortCode)

	response := types.ShortenResponse{
		ShortURL:    targetURL,
		OriginalURL: originalURL.Url,
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(response); err != nil {
		http.Error(res, "failed to encode response", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(buf.Bytes())
}
