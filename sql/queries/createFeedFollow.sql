-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (created_at, updated_at, user_id, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *
)
SELECT 
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users
ON inserted_feed_follow.user_id = users.id
INNER JOIN feeds
ON inserted_feed_follow.feed_id = feeds.id;


/*CREATE TABLE feed_follows (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID UNIQUE NOT NULL REFERENCES users (id)
        ON DELETE CASCADE,
    feed_id SERIAL UNIQUE NOT NULL REFERENCES feeds (id)
        ON DELETE CASCADE
);*/