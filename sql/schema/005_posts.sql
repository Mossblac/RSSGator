-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    title TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    description TEXT,
    published_at TIMESTAMP,
    feed_id INTEGER NOT NULL REFERENCES feeds (id)
        ON DELETE CASCADE,
    UNIQUE (feed_id);
)

-- +goose Down
DROP TABLE posts;