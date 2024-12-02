package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var conn *pgx.Conn

func main() {
	var err error
	if err := godotenv.Load(); err != nil {
		fmt.Println("failed to load .env file")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	conn, err = pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var name string
	var value float64

	conn.QueryRow(context.Background(), "Select name, value from playing_with_neon where id=$1", 2).Scan(&name, &value)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name, value)
}
