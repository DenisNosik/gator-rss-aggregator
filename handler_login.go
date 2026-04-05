package main

import (
	"context"
	"errors"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("Login command expects a username")
	}

	user, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return err
	}

	return nil
}
