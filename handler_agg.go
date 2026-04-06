package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/DenisNosik/gator-rss-aggregator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return errors.New("Agg command expects collecting feeds time e.g(1m, 5m, 1h etc.)")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenReqs)
	fmt.Println()

	ticker := time.NewTicker(timeBetweenReqs)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	feedToFetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	if err := s.db.MarkFeedFetched(context.Background(), feedToFetch.ID); err != nil {
		return err
	}

	rssFeed, err := fetchFeed(context.Background(), feedToFetch.Url)
	if err != nil {
		return err
	}

	for _, rssItem := range rssFeed.Channel.Item {
		pubDate, err := time.Parse(time.RFC1123Z, rssItem.PubDate)
		if err != nil {
			pubDate, err = time.Parse(time.RFC1123, rssItem.PubDate)
		}
		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       rssItem.Title,
			Url:         rssItem.Link,
			Description: sql.NullString{String: rssItem.Description, Valid: rssItem.Description != ""},
			PublishedAt: sql.NullTime{Time: pubDate, Valid: err == nil},
			FeedID:      feedToFetch.ID,
		})
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
				continue
			}
			return err
		}
	}

	return nil
}
