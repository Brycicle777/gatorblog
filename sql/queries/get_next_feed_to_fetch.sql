-- name: GetNextFeedToFetch :one
SELECT
          id
          ,created_at
          ,updated_at
          ,last_fetched_at
          ,name
          ,url
FROM      feeds
ORDER BY  last_fetched_at ASC NULLS FIRST
LIMIT 1;