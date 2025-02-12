-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
  VALUES (
      $1,
      $2,
      $3,
      $4,
      $5
  )
  RETURNING *
)
SELECT
      inserted_feed_follow.*
      ,feeds.name AS feedName
      ,users.name AS userName
FROM  inserted_feed_follow
INNER JOIN users ON users.ID = inserted_feed_follow.user_id
INNER JOIN feeds ON feeds.ID = inserted_feed_follow.feed_id;