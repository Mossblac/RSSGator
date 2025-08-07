-- name: GetFeeds :many
SELECT f.*, u.name AS created_by
FROM feeds f
JOIN users u ON f.user_id = u.id;