package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/DenisNosik/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("Login command expects a username")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	})
	if err != nil {
		return err
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return err
	}

	s.cfg.CurrentUserName = cmd.Args[0]
	fmt.Printf("User: %s was created.\n", user.Name)

	return nil
}
