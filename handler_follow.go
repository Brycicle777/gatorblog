package main

import (
	"context"
	"fmt"
	"time"

	"internal/database"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
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

func handlerFollowing(s *state, cmd command, user database.User) error {
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

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}
	url := cmd.Args[0]
	_, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error retrieving feed details, it may not exist: %v", url)
	}

	err = s.db.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		Name: user.Name,
		Url:  url,
	})
	if err != nil {
		return fmt.Errorf("error unfollowing feed: %v", err)
	}
	fmt.Println("Successfully unfollowed feed.")
	return nil
}
