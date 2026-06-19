package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func Connect() (*pgx.Conn, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		user, password, host, port, dbName,
	)

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
