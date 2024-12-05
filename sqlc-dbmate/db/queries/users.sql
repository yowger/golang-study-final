-- name: GetUserByEmail :one
SELECT * FROM dbmate WHERE email = $1;

-- name: CreateUser :exec
INSERT INTO dbmate (name, email) VALUES ($1, $2);
