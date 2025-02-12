// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: unfollow_feed.sql

package database

import (
	"context"
)

const unfollowFeed = `-- name: UnfollowFeed :exec
DELETE
FROM    feed_follows
USING   users, feeds
WHERE   users.name = $1
    AND feeds.url = $2
`

type UnfollowFeedParams struct {
	Name string
	Url  string
}

func (q *Queries) UnfollowFeed(ctx context.Context, arg UnfollowFeedParams) error {
	_, err := q.db.ExecContext(ctx, unfollowFeed, arg.Name, arg.Url)
	return err
}
