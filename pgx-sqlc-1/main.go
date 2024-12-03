package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"pgx-sqlc-1/internal/db"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("failed to load .env file")
	}

	dbUrl := os.Getenv("DATABASE_URL")

	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		log.Fatalf("Failed to parse database URL: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()

	queries := db.New(pool)

	createdGamer, err := queries.CreateGamer(context.Background(), db.CreateGamerParams{
		FirstName: "Tiago",
		LastName:  "Taron",
	})
	if err != nil {
		log.Fatalf("Failed to create player: %v", err)
	}
	fmt.Printf("Created player with ID: %d\n", createdGamer.ID)

	gamers, err := queries.GetGamers(context.Background())
	if err != nil {
		log.Fatalf("Failed to get players: %v", err)
	}

	for _, gamer := range gamers {
		fmt.Printf("Player ID: %d, Name: %s %s\n", gamer.ID, gamer.FirstName, gamer.LastName)
	}
}

// func getUsers(pool *pgxpool.Pool) {
// 	ctx := context.Background()
// 	rows, err := pool.Query(ctx, "SELECT id, first_name, last_name FROM users")
// 	if err != nil {
// 		log.Fatalf("Query failed: %v", err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var (
// 			id         int
// 			first_name string
// 			last_name  string
// 		)
// 		if err := rows.Scan(&id, &first_name, &last_name); err != nil {
// 			log.Fatalf("Row scan failed: %v", err)
// 		}
// 		log.Printf("User ID: %d, first name: %s, last name: %s", id, first_name, last_name)
// 	}

// 	if rows.Err() != nil {
// 		log.Fatalf("Rows iteration failed: %v", rows.Err())
// 	}
// }
