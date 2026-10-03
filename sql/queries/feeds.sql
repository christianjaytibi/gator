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

-- name: GetFeedByUrl :one
SELECT 
  *
FROM
  feeds
WHERE url = $1; 

-- name: MarkFeedFetch :one
UPDATE feeds
SET 
  last_fetched_at = NOW(),
  updated_at = NOW()
WHERE 
  id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
ORDER BY last_fetched_at NULLS FIRST
LIMIT 1;
