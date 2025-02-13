package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"internal/database"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("User-Agent", "gator")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %v", err)
	}
	defer res.Body.Close()

	xmlData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading data: %v", err)
	}

	var rssFeed RSSFeed
	if err := xml.Unmarshal(xmlData, &rssFeed); err != nil {
		return nil, fmt.Errorf("error unmarshalling xml: %v", err)
	}

	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for i := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[i].Title = html.UnescapeString(rssFeed.Channel.Item[i].Title)
		rssFeed.Channel.Item[i].Description = html.UnescapeString(rssFeed.Channel.Item[i].Description)
	}

	return &rssFeed, nil
}

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("error getting next feed to fetch: %v", err)
	}
	_, err = s.db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %v", err)
	}
	rssFeed, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed: %v", err)
	}

	for i := range rssFeed.Channel.Item {
		valid := true
		if len(rssFeed.Channel.Item[i].Description) == 0 {
			valid = false
		}
		nullDesc := sql.NullString{
			String: rssFeed.Channel.Item[i].Description,
			Valid:  valid,
		}
		pubDate, err := time.Parse(time.RFC1123Z, rssFeed.Channel.Item[i].PubDate)
		if err != nil {
			log.Printf("error parsing pubdate %v: %v\n", rssFeed.Channel.Item[i].PubDate, err)
			continue
		}
		existingPost, err := s.db.GetPostByUrl(ctx, rssFeed.Channel.Item[i].Link)
		if err != nil && err != sql.ErrNoRows {
			log.Printf("error checking if url %v exists already: %v", rssFeed.Channel.Item[i].Link, err)
			continue
		}
		if len(existingPost.Url) != 0 {
			// log.Printf("post %v exists, skipping...", existingPost.Url)
			continue
		}
		post, err := s.db.CreatePost(ctx, database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       rssFeed.Channel.Item[i].Title,
			Url:         rssFeed.Channel.Item[i].Link,
			Description: nullDesc,
			PublishedAt: pubDate,
			FeedID:      feed.ID,
		})
		if err != nil {
			log.Printf("error creating post %v: %v\n", rssFeed.Channel.Item[i].Link, err)
			continue
		}

		fmt.Printf("Created post for: %v\n", post.Url)
	}
	return nil
}
