package ext

import (
	"context"
	"fmt"
	"time"

	"github.com/Mossblac/RSSGator/internal/database"
)

func HandlerAddFeed(s *State, cmd Command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("not enough arguments provided")
	}
	name := cmd.Args[0]
	url := cmd.Args[1]
	userId, err := s.DataBase.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("unable to obtain user UUID: %v", err)
	}

	params := database.CreateFeedParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    userId.ID,
	}

	outputFeed, err := s.DataBase.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to create feed: %v", err)
	}

	fmt.Printf("Feed created:\nName: %s\nURL: %s\nUserID: %s\n", outputFeed.Name, outputFeed.Url, outputFeed.UserID)
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
