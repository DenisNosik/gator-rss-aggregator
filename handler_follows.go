package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/DenisNosik/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("Follow command expects a url")
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%s successfully followed to the %s\n", feedFollow.UserName, feedFollow.FeedName)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	if len(feedFollows) == 0 {
		fmt.Println("You don't follow any feed")
		return nil
	}

	fmt.Println("You following: ")
	for _, ff := range feedFollows {
		fmt.Println(ff.FeedName)
	}

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("Unfollow command expects a url")
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}

	if err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}); err != nil {
		return err
	}

	fmt.Printf("%s successfully unfollowed from %s\n", user.Name, feed.Name)

	return nil
}
