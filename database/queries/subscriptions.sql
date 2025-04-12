-- name: CreateSubscription :one
INSERT INTO subscriptions (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteSubscription :execrows
DELETE FROM subscriptions WHERE user_id = $1 AND feed_id = $2;

-- name: GetActiveSubscriptions :many
SELECT
    s.*,
    sqlc.embed(f)
FROM
    subscriptions s
    INNER JOIN feeds f ON f.id = s.feed_id
WHERE
    s.user_id = $1;
