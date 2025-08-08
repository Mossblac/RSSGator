-- name: DeleteAllUsers :exec
DELETE FROM users;

-- name: ResetFeedsSequence :exec
SELECT setval('feeds_id_seq', 1, false);

-- name: ResetFeedFollowsSequence :exec
SELECT setval('feed_follows_id_seq', 1, false);

-- name: ResetPosts :exec
SELECT setval('posts_id_seq', 1, false);