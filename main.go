package main

import (
	"fmt"

	"github.com/Mossblac/RSSGator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		return
	}

	err = cfg.SetUser("Mossblac")
	if err != nil {
		fmt.Printf("Error setting user: %v\n", err)
		return
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("Error reading config again: %v\n", err)
		return
	}

	fmt.Printf("Config: %+v\n", cfg)

}
