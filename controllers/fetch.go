package controllers

import (
	"context"
	"net/http"
	"urlshort/internal/db"

	"github.com/jackc/pgx/v5"
)

func Fetch(res http.ResponseWriter, req *http.Request) {
	code := req.PathValue("code")

	var originalUrl string
	err := db.DB.QueryRow(
		context.Background(),
		`
		UPDATE urls
		SET click_count = click_count + 1
		WHERE short_code = $1
		RETURNING original_url;
		`,
		code,
	).Scan(&originalUrl)

	if err == pgx.ErrNoRows {
		http.Error(res, "shortened URL doesn't exist", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(res, "db error", http.StatusInternalServerError)
		return
	}

	http.Redirect(res, req, originalUrl, 302)
}
