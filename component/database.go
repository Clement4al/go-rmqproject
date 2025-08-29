package component

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

var Database *pgx.Conn

func Startdb() {
	// Connection string
	dsn := "postgres://postgres:test1234@localhost:5432/postgres"

	// Connection to PostgreSQL
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// Save connection in the global variable
	Database = conn

	fmt.Println("Database connection created successfully..")
}

// CloseDB closes the global database connection
func CloseDB() {
	if Database != nil {
		err := Database.Close(context.Background())
		if err != nil {
			log.Printf("Error closing database: %v\n", err)
		} else {
			fmt.Println("Database connection closed successfully..")
		}
	}
}
