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
	// get env variables
	var err error
	if err := godotenv.Load(); err != nil {
		fmt.Println("failed to load .env file")
	}

	// connect db
	dbUrl := os.Getenv("DATABASE_URL")
	conn, err = pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// create user
	newUser := User{Name: "John Doe", Value: 1.0}
	userID, err := createUser(context.Background(), newUser)
	if err != nil {
		fmt.Fprintf(os.Stderr, "createUser: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Created user with ID: ", userID)

	// update user
	updateUser := User{ID: 1, Name: "Jane Doe", Value: 2.0}
	updateErr := updateUserByID(context.Background(), updateUser)
	if updateErr != nil {
		fmt.Fprintf(os.Stderr, "updateUserByID: %v\n", err)
		os.Exit(1)
	}

	// get all users
	users, err := getAllUsers(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "getAllUsers: %v\n", err)
		os.Exit(1)
	}
	// print all users
	for _, user := range users {
		fmt.Println("User: ", user)
	}

	// delete user
	deleteErr := deleteUser(context.Background(), 1)
	if deleteErr != nil {
		fmt.Fprintf(os.Stderr, "deleteUser: %v\n", err)
		os.Exit(1)
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

func updateUserByID(ctx context.Context, user User) error {
	commandTag, err := conn.Exec(ctx, "UPDATE playing_with_neon SET name = $1, value = $2 WHERE id = $3", user.Name, user.Value, user.ID)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil

}

func deleteUser(ctx context.Context, id int) error {
	commandTag, err := conn.Exec(ctx, "DELETE FROM users WHERE id = $1", id)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no user found with id %d", id)
	}

	return nil
}

/*
	conn.QueryRow(context.Background(), "Select name, value from playing_with_neon where id=$1", 2).Scan(&name, &value)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
*/
