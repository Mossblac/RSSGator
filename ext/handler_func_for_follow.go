package ext

import (
	"context"
	"fmt"
	"time"

	"github.com/Mossblac/RSSGator/internal/database"
)

func HandlerFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("no url provided")
	}

	inputURL := cmd.Args[0]

	feed, err := s.DataBase.FindFeedFromUrl(context.Background(), inputURL)
	if err != nil {
		return fmt.Errorf("unable to find feed that matches url: %v", err)
	}

	params := database.CreateFeedFollowParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	newFollow, err := s.DataBase.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to create follow entry: %v", err)
	}

	fmt.Printf("User:%v\nNow Following:%v\n", s.Config.CurrentUserName, newFollow.FeedName)

	return nil

}

/*type CreateFeedFollowParams struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    int32
}

type CreateFeedFollowRow struct {
	ID        int32
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    int32
	FeedName  string
	UserName  string
}*/

/*func (q *Queries) FindFeedFromUrl(ctx context.Context, url string) (Feed, error)
func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (CreateFeedFollowRow, error)*/
