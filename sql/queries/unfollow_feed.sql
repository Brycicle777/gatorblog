-- name: UnfollowFeed :exec
DELETE
FROM    feed_follows
USING   users, feeds
WHERE   users.name = $1
    AND feeds.url = $2;