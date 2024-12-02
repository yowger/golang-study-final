package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var conn *pgx.Conn

type User struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Value float64 `json:"email"`
}

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

	newUser := User{Name: "John Doe", Value: 1.0}
	userID, err := createUser(context.Background(), newUser)
	if err != nil {
		fmt.Fprintf(os.Stderr, "createUser: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Created user with ID: ", userID)

	users, err := getAllUsers(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "getAllUsers: %v\n", err)
		os.Exit(1)
	}

	for _, user := range users {
		fmt.Println("User: ", user)
	}

}

func createUser(ctx context.Context, user User) (int, error) {
	var id int
	err := conn.QueryRow(ctx, "INSERT INTO playing_with_neon (name, value) VALUES ($1,$2) RETURNING id", user.Name, user.Value).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func getAllUsers(ctx context.Context) ([]User, error) {
	rows, err := conn.Query(ctx, "SELECT id, name, value from playing_with_neon")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Value)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

/*
	conn.QueryRow(context.Background(), "Select name, value from playing_with_neon where id=$1", 2).Scan(&name, &value)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
*/
