-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, user_id)
VALUES
  ($1, $2, $3, $4)
RETURNING *;

-- name: ListFeeds :many
SELECT
  f.id,
  f.name,
  f.url,
  u.name AS username
FROM 
  feeds f
JOIN
  users u
ON 
  f.user_id = u.id;