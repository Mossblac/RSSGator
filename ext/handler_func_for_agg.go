package ext

import (
	"fmt"
	"time"
)

func HandlerAgg(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("no duration set")
	}
	time_between_reqs := cmd.Args[0]

	duration, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return fmt.Errorf("error setting time duration: %v", err)
	}

	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		ScrapeFeeds(s)
		fmt.Printf("Collecting feeds every %v", duration)

	}
}
