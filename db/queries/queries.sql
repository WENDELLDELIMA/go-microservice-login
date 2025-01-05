-- name: CreateUser :one
INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)
RETURNING id, username, email, created_at;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;
