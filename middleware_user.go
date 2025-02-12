package main

import (
	"context"
	"fmt"
	"internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("error retrieving current user: %v", s.cfg.CurrentUserName)
		}
		return handler(s, cmd, user)
	}
}
