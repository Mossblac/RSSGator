-- name: GetFeedFollowForUser :many
SELECT ff.*, f.name AS feed_name, u.name AS user_name
FROM feed_follows ff
JOIN users u ON ff.user_id = u.id
JOIN feeds f ON f.id= ff.feed_id
WHERE u.id = $1;

