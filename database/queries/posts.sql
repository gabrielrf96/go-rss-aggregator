-- name: CreatePost :one
INSERT INTO posts (
    id,
    created_at,
    updated_at,
    published_at,
    title,
    description,
    url,
    feed_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetPostsForUser :many
SELECT
    p.*,
    sqlc.embed(f)
FROM
    subscriptions s
    INNER JOIN posts p ON p.feed_id = s.feed_id
    INNER JOIN feeds f ON f.id = s.feed_id
WHERE
    s.user_id = $1
ORDER BY
    p.published_at DESC
LIMIT $2;
