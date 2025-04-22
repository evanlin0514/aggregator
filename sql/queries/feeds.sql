-- name: Addfeed :one
INSERT INTO feeds (name, url, user_id)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedByUrl :one
SELECT id, name FROM feeds
WHERE feeds.url = $1;

-- name: CreateFeedFollow :one
WITH insert_feed_follow AS (
    INSERT INTO feed_follows (user_id, feed_id)
    VALUES(
        $1,
        $2
    )
    RETURNING *
) SELECT
    insert_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM insert_feed_follow
INNER JOIN feeds ON insert_feed_follow.feed_id = feeds.id
INNER JOIN users ON insert_feed_follow.user_id = users.id;

-- name: GetFeedFollowsForUser :many
SELECT 
    f.name AS feed_name, 
    u.name AS user_name, 
    ff.created_at, 
    ff.updated_at
FROM feed_follows ff
INNER JOIN feeds f ON ff.feed_id = f.id
INNER JOIN users u ON ff.user_id = u.id
WHERE ff.user_id = $1;