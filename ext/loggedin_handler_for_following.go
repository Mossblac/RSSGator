package ext

import (
	"context"
	"fmt"

	"github.com/Mossblac/RSSGator/internal/database"
)

func HandlerFollowing(s *State, cmd Command, user database.User) error {
	user, err := s.DataBase.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("no user current user found: %v", err)
	}

	feedFollows, err := s.DataBase.GetFeedFollowForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("unable to obtain follows list: %v", err)
	}

	for _, feed := range feedFollows {
		fmt.Printf("Feed:%v\nUser:%v\n", feed.FeedName, feed.UserName)
	}

	return nil
}

//func (q *Queries) GetFeedFollowForUser(ctx context.Context, userID uuid.UUID) ([]FeedFollow, error)
//func (q *Queries) GetUser(ctx context.Context, name string) (User, error)
