-- name: CreateFeed :exec
CREATE TABLE feeds (
    id TEXT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT,
    url TEXT NOT NULL,
    user_id TEXT NOT NULL,
    CONSTRAINT fk_user_id
        FOREIGN KEY (user_id)
        REFERENCES users (id)
        ON DELETE CASCADE
);