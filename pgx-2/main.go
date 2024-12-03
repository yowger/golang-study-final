package main

import (
	"context"
	"fmt"
	"log"
	"os"

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

	getUsers(pool)
}

func getUsers(pool *pgxpool.Pool) {
	ctx := context.Background()
	rows, err := pool.Query(ctx, "SELECT id, first_name, last_name FROM users")
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id         int
			first_name string
			last_name  string
		)
		if err := rows.Scan(&id, &first_name, &last_name); err != nil {
			log.Fatalf("Row scan failed: %v", err)
		}
		log.Printf("User ID: %d, first name: %s, last name: %s", id, first_name, last_name)
	}

	if rows.Err() != nil {
		log.Fatalf("Rows iteration failed: %v", rows.Err())
	}
}

func createUser(pool *pgxpool.Pool, name string) {
	ctx := context.Background()
	commandTag, err := pool.Exec(ctx, "INSERT INTO users (name) VALUES ($1)", name)
	if err != nil {
		log.Fatalf("Insert failed: %v", err)
	}
	log.Printf("Rows affected: %d", commandTag.RowsAffected())
}

func updateUser(pool *pgxpool.Pool, id int, name string) {
	ctx := context.Background()
	commandTag, err := pool.Exec(ctx, "UPDATE users SET name = $1 WHERE id = $2", name, id)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}
	log.Printf("Rows affected: %d", commandTag.RowsAffected())
}

/*
	advance
		prepared statements
		transactions
		timeouts

	func usePreparedStatements(pool *pgxpool.Pool) {
		ctx := context.Background()
		stmtName := "getUserByID"

		// Prepare the statement
		_, err := pool.Prepare(ctx, stmtName, "SELECT name FROM users WHERE id=$1")
		if err != nil {
			log.Fatalf("Failed to prepare statement: %v", err)
		}

		var name string
		err = pool.QueryRow(ctx, stmtName, 1).Scan(&name)
		if err != nil {
			log.Fatalf("QueryRow failed: %v", err)
		}
		log.Printf("Name: %s", name)
	}

	func runTransaction(pool *pgxpool.Pool) {
		ctx := context.Background()
		tx, err := pool.Begin(ctx)
		if err != nil {
			log.Fatalf("Failed to begin transaction: %v", err)
		}
		defer tx.Rollback(ctx) // Ensure rollback in case of failure

		_, err = tx.Exec(ctx, "INSERT INTO users (name) VALUES ($1)", "Alice")
		if err != nil {
			log.Fatalf("Transaction step failed: %v", err)
		}

		_, err = tx.Exec(ctx, "INSERT INTO users (name) VALUES ($1)", "Bob")
		if err != nil {
			log.Fatalf("Transaction step failed: %v", err)
		}

		// Commit the transaction
		if err := tx.Commit(ctx); err != nil {
			log.Fatalf("Failed to commit transaction: %v", err)
		}
		log.Println("Transaction committed successfully!")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := pool.QueryRow(ctx, "SELECT id FROM users WHERE name=$1", "Alice")
	var id int
	if err := row.Scan(&id); err != nil {
		log.Printf("Query failed or context canceled: %v", err)
	}
*/
