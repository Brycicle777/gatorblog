package main

import (
	"context"
	"fmt"
	"time"

	"internal/database"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error retrieving current user: %v", s.cfg.CurrentUserName)
	}
	url := cmd.Args[0]
	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error retrieving feed id for url: %v", url)
	}

	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error following feed: %v", err)
	}
	fmt.Printf("user %v now following feed %v\n", follow.Username, follow.Feedname)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error retrieving current user: %v", s.cfg.CurrentUserName)
	}

	feed_follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return fmt.Errorf("error retrieving follows for user %v: %v", user.Name, err)
	}
	fmt.Println("current user is following these feeds:")
	for i := range feed_follows {
		fmt.Printf("%v\n", feed_follows[i].Feedname)
	}

	return nil
}
