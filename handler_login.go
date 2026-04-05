package main

import "errors"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("Login command expects a username")
	}

	if err := s.cfg.SetUser(cmd.Args[0]); err != nil {
		return err
	}

	return nil
}
