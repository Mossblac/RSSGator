package ext

import (
	"context"
	"fmt"
)

func HandlerAgg(s *State, cmd Command) error {
	rss, err := FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("fetch feed error: %v", err)
	}

	fmt.Printf("%+v\n", rss)

	return nil
}
