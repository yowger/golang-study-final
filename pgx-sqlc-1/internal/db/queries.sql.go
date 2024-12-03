// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package db

import (
	"context"
)

const createGamer = `-- name: CreateGamer :one
INSERT INTO gamers (first_name, last_name)
VALUES ($1, $2)
RETURNING id, first_name, last_name
`

type CreateGamerParams struct {
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
}

func (q *Queries) CreateGamer(ctx context.Context, arg CreateGamerParams) (Gamer, error) {
	row := q.db.QueryRow(ctx, createGamer, arg.FirstName, arg.LastName)
	var i Gamer
	err := row.Scan(&i.ID, &i.FirstName, &i.LastName)
	return i, err
}

const createTodo = `-- name: CreateTodo :one
INSERT INTO todos (user_id, task, done)
VALUES ($1, $2, $3)
RETURNING id, user_id, task, done
`

type CreateTodoParams struct {
	UserID int32  `db:"user_id" json:"user_id"`
	Task   string `db:"task" json:"task"`
	Done   bool   `db:"done" json:"done"`
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.db.QueryRow(ctx, createTodo, arg.UserID, arg.Task, arg.Done)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Task,
		&i.Done,
	)
	return i, err
}

const getGamer = `-- name: GetGamer :one
SELECT id, first_name, last_name
FROM gamers
WHERE id = $1
`

func (q *Queries) GetGamer(ctx context.Context, id int32) (Gamer, error) {
	row := q.db.QueryRow(ctx, getGamer, id)
	var i Gamer
	err := row.Scan(&i.ID, &i.FirstName, &i.LastName)
	return i, err
}

const getGamers = `-- name: GetGamers :many
SELECT id, first_name, last_name
FROM gamers
`

func (q *Queries) GetGamers(ctx context.Context) ([]Gamer, error) {
	rows, err := q.db.Query(ctx, getGamers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Gamer
	for rows.Next() {
		var i Gamer
		if err := rows.Scan(&i.ID, &i.FirstName, &i.LastName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const deleteGamer = `-- name: deleteGamer :exec
DELETE FROM gamers
Where id = $1
`

func (q *Queries) deleteGamer(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteGamer, id)
	return err
}
