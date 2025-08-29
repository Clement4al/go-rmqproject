package main

import (
	"context"
	"fmt"
	"log"

	"Go/component"
)

func main() {
	// component.StartRabbitMQ()
	component.Startdb()
	defer component.CloseDB()
	//exmaple query
	var sweepapp string
	err := component.Database.QueryRow(context.Background(), "SELECT 'hello from the Database'").Scan(&sweepapp)
	if err != nil {
		log.Fatalf("Query failed: %v\n", err)
	}

	fmt.Println(sweepapp)
}
