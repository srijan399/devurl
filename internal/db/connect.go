package db

import (
	"context"
	"fmt"
	"log"
)

func ConnectMain() {
	conn, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	DB = conn

	fmt.Println("Connected to Postgres!")

	if err := createTables(); err != nil {
		log.Fatal("Failed to create tables:", err)
	}
}

func createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS urls (
		id SERIAL PRIMARY KEY,
		short_code TEXT UNIQUE NOT NULL,
		original_url TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	);`

	_, err := DB.Exec(context.Background(), query)
	return err
}
