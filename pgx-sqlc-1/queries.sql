-- name: GetGamers :many
SELECT *
FROM gamers;
-- name: GetGamer :one
SELECT *
FROM gamers
WHERE id = $1;
-- name: deleteGamer :exec
DELETE FROM gamers
Where id = $1;
-- name: CreateGamer :one
INSERT INTO gamers (first_name, last_name)
VALUES ($1, $2)
RETURNING *;
-- name: CreateTodo :one
INSERT INTO todos (user_id, task, done)
VALUES ($1, $2, $3)
RETURNING *;