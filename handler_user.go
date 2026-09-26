package main

import (
	"context"
	"fmt"
	"log"

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

func handleReset(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	ctx := context.Background()
	if err := s.db.DeleteAllAuthors(ctx); err != nil {
		log.Fatalf("Reset unsuccessful: %v", err)
	}

	return nil
}

func handleListUserNames(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	ctx := context.Background()
	usernames, err := s.db.ListUserNames(ctx)
	if err != nil {
		return err
	}

	for _, name := range usernames {
		fmt.Printf("* %s ", name)
		if name == s.cfg.CurrentUserName {
			fmt.Printf("(current)")
		}

		fmt.Printf("\n")
	}

	return nil
}
