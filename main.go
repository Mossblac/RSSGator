package main

import (
	"fmt"
	"os"

	"github.com/Mossblac/RSSGator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("no arguments provided")
		os.Exit(1)
	}

	commandName := os.Args[1]
	argSlice := os.Args[1:]

}
