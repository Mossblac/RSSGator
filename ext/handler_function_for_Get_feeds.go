package ext

import (
	"context"
	"fmt"
)

func HandlerGetFeeds(s *State, cmd Command) error {
	feeds, err := s.DataBase.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("unable to read feeds table: %v", err)
	}

	for _, feed := range feeds {
		fmt.Printf("FeedName: %v\nCreated: %v\nUpdated: %v\nFeedURL: %v\nUserID: %v\nCreatedBy: %v\n\n", feed.Name, feed.CreatedAt, feed.UpdatedAt, feed.Url, feed.UserID, feed.CreatedBy)
	}
	return nil
}

/*type GetFeedsRow struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Url       string
	UserID    uuid.UUID
	CreatedBy string
}*/
