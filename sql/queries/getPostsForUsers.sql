-- name: GetPostsForUser :many
SELECT * FROM posts 
ORDER BY published_at ASC NULLS FIRST
LIMIT $1;