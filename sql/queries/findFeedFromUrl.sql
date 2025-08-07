-- name: FindFeedFromUrl :one
SELECT * FROM feeds WHERE url = $1;