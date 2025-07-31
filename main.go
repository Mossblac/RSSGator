package main

import (
	"fmt"
	"os"

	"github.com/Mossblac/RSSGator/ext"
	"github.com/Mossblac/RSSGator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config: %v", err)
		os.Exit(1)
	}

	currentState := ext.State{
		Config: &cfg,
	}

	handler := make(map[string]func(*ext.State, ext.Command) error)

	activeCommands := ext.Commands{
		Handlers: handler,
	}

	activeCommands.Register("login", ext.HandlerLogin)

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
		fmt.Printf("error running commands: %v", err)
		os.Exit(1)
	}

}
