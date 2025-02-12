package main

import (
	"context"
	"fmt"
	"time"

	"internal/database"

	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	//wagslan can't be reached on work wifi
	url := "https://www.wagslane.dev/index.xml"
	ctx := context.Background()

	rssFeed, err := fetchFeed(ctx, url)
	if err != nil {
		return fmt.Errorf("error retrieving from url: %v", err)
	}

	fmt.Printf("%v", rssFeed)
	return nil
}

func handlerAddfeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed: %v", err)
	}

	fmt.Printf("Feed created:\n%v\n", feed)

	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error following new feed: %v", err)
	}
	fmt.Printf("User %v now following new feed %v\n", follow.Username, follow.Feedname)

	return nil

}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error retrieving all feeds: %v", err)
	}
	for i := range feeds {
		fmt.Printf("%v\n", feeds[i])
	}
	return nil
}
