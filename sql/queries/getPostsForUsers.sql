-- name: GetPostsForUser :many
SELECT posts.*, feeds.name as feed_name
FROM posts 
JOIN feeds ON posts.feed_id = feeds.id
WHERE feeds.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;