package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/christianjaytibi/gator/internal/config"
	"github.com/christianjaytibi/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	appState := state{
		db:  dbQueries,
		cfg: &cfg,
	}

	appCommands := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	appCommands.register("login", handlerLogin)
	appCommands.register("register", handleRegister)
	appCommands.register("reset", handleReset)
	appCommands.register("users", handleListUserNames)
	appCommands.register("agg", handlerAgg)
	appCommands.register("addfeed", handleAddFeed)
	appCommands.register("feeds", handleListFeeds)

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
