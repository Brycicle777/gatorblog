-- name: GetPostByUrl :one
SELECT
            posts.id
            ,posts.created_at
            ,posts.updated_at
            ,posts.title
            ,posts.url
            ,posts.description
            ,posts.published_at
            ,posts.feed_id
FROM        posts
WHERE       posts.url = $1;
