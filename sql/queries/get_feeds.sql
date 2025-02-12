-- name: GetFeeds :many
SELECT
            feeds.id
            ,feeds.created_at
            ,feeds.updated_at
            ,feeds.name AS feedName
            ,feeds.url
            ,users.name AS userName
FROM        feeds
INNER JOIN  users ON users.ID = feeds.user_id;