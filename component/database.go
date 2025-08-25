package component

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func Startdb() {
	// Connection string
	dsn := "postgres://postgres:test1234@localhost:5432/postgres"

	// Connect to PostgreSQL
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	// Test query
	var greeting string
	err = conn.QueryRow(context.Background(), "SELECT 'Hello, PostgreSQL from Go!'").Scan(&greeting)
	if err != nil {
		log.Fatalf("Query failed: %v\n", err)
	}

	fmt.Println(greeting)
}
