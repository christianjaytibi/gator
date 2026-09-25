package main

import (
	"log"
	"os"

	"github.com/christianjaytibi/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	appState := state{
		cfg: &cfg,
	}

	appCommands := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	appCommands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: gator <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	err = appCommands.run(&appState, command{
		Name: cmdName,
		Args: cmdArgs,
	})

	if err != nil {
		log.Fatal(err)
	}
}
