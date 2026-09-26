-- name: CreateUser :one
INSERT INTO users (id, name)
VALUES 
  ($1, $2)
RETURNING *;

-- name: GetUser :one
SELECT * 
FROM users
WHERE
  name = $1;

-- name: DeleteAllAuthors :exec
DELETE FROM users;

-- name: ListUserNames :many
SELECT name
FROM users;