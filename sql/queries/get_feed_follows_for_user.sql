-- name: GetFeedFollowsForUser :many
SELECT
            feed_follows.ID
            ,feed_follows.created_at
            ,feed_follows.created_at
            ,feeds.name AS feedName
            ,users.name AS userName
FROM        feed_follows
INNER JOIN  feeds ON feeds.ID = feed_follows.feed_id
INNER JOIN  users ON users.ID = feed_follows.user_id
WHERE       users.name = $1;