package main

import (
	"context"
	"fmt"
	"time"

	"internal/database"

	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <time_duration>", cmd.Name)
	}
	time_between_reqs := cmd.Args[0]
	dur, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return fmt.Errorf("error parsing duration string %v with error: %v", time_between_reqs, err)
	}
	ticker := time.NewTicker(dur)
	fmt.Printf("Collecting feeds every %v\n", time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
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
