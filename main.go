package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Mossblac/RSSGator/ext"
	"github.com/Mossblac/RSSGator/internal/config"
	"github.com/Mossblac/RSSGator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config: %v", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		fmt.Printf("error opening database: %v", err)
	}

	dbQueries := database.New(db)

	currentState := ext.State{
		DataBase: dbQueries,
		Config:   &cfg,
	}

	handler := make(map[string]func(*ext.State, ext.Command) error)

	activeCommands := ext.Commands{
		Handlers: handler,
	}

	activeCommands.Register("login", ext.HandlerLogin)
	activeCommands.Register("register", ext.HandlerRegister)
	activeCommands.Register("reset", ext.HandlerReset)
	activeCommands.Register("users", ext.HandlerUsers)
	activeCommands.Register("agg", ext.HandlerAgg)
	activeCommands.Register("addfeed", ext.HandlerAddFeed)

	if len(os.Args) < 2 {
		fmt.Println("no arguments")
		os.Exit(1)
	}

	commandPrompt := os.Args[1]
	argumentsString := os.Args[2:]

	inputCommand := ext.Command{
		CommandName: commandPrompt,
		Args:        argumentsString,
	}

	err = activeCommands.Run(&currentState, inputCommand)
	if err != nil {
		fmt.Printf("error running commands: %v\n", err)
		os.Exit(1)
	}

}
