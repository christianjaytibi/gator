package main

import (
	"context"
	"fmt"

	"github.com/christianjaytibi/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	ctx := context.Background()
	username := cmd.Args[0]
	user, err := s.db.GetUser(ctx, username)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("user switched!")
	return nil
}

func handleRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	ctx := context.Background()
	username := cmd.Args[0]

	user, err := s.db.CreateUser(ctx, database.CreateUserParams{
		ID:   uuid.New(),
		Name: username,
	})
	if err != nil {
		return err
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Println("user created!")
	return nil
}
