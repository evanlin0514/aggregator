-- name: CreateUser :one
INSERT INTO users (id, create_at, update_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: CheckUser :one
SELECT * FROM users
WHERE name = $1;

-- name: Reset :exec
DELETE FROM users;

-- name: GetUser :many
SELECT name FROM users;

-- name: GetName :one
SELECT name FROM users
WHERE id = $1;

-- name: GetId :one
SELECT id FROM users
WHERE name = $1;

