package ext

import (
	"context"
	"fmt"
	"time"

	"github.com/Mossblac/RSSGator/internal/database"
)

func HandlerAddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("not enough arguments provided")
	}
	name := cmd.Args[0]
	url := cmd.Args[1]

	params := database.CreateFeedParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	outputFeed, err := s.DataBase.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to create feed: %v", err)
	}

	fmt.Printf("Feed added to table:\nName: %s\nURL: %s\nCreated By: %s\n", outputFeed.Name, outputFeed.Url, user.Name)

	paramsForFollow := database.CreateFeedFollowParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    outputFeed.ID,
	}
	newFollow, err := s.DataBase.CreateFeedFollow(context.Background(), paramsForFollow)
	if err != nil {
		return fmt.Errorf("failed to create follow entry: %v", err)
	}

	fmt.Printf("User:%v\nNow Following:%v\n", s.Config.CurrentUserName, newFollow.FeedName)

	return nil
}

/*type CreateFeedParams struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Url       string
	UserID    uuid.UUID
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error)*/
