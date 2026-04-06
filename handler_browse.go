package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/DenisNosik/gator-rss-aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32 = 2
	if len(cmd.Args) != 0 {
		parsed, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid limit: %w", err)
		}
		limit = int32(parsed)
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return err
	}

	for _, post := range posts {
		if post.PublishedAt.Valid {
			fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		}
		fmt.Printf("* %s\n", post.Title)
		if post.Description.Valid {
			fmt.Printf("* %v\n", post.Description.String)
		}
		fmt.Printf("* Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}
