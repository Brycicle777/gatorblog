package main

import (
	"context"
	"fmt"
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
