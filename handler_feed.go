package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/DenisNosik/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return errors.New("Add Feed command expects a feed name and url")
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    currentUser.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println(feed.Name)
	fmt.Println("Feed created successfully")

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))

	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Println("==========")
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println(user.Name)
		fmt.Println("==========")
	}
	return nil
}
