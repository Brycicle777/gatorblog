-- name: GetFeed :one
SELECT
        id
        ,created_at
        ,updated_at
        ,name
        ,url
FROM    feeds
WHERE   url = $1;