package ext

import (
	"context"
	"fmt"

	"github.com/Mossblac/RSSGator/internal/database"
)

func HandlerUnfollow(s *State, cmd Command, user database.User) error {
	params := database.FindToUnfollowParams{
		UserID: user.ID,
		Url:    cmd.Args[0],
	}

	err := s.DataBase.FindToUnfollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error unfollowing: %v", err)
	}
	fmt.Printf("%v\nUnfollowed:%v\n", s.Config.CurrentUserName, cmd.Args[0])
	return nil
}

/*type FindToUnfollowParams struct {
	UserID uuid.UUID
	Url    string
}

func (q *Queries) FindToUnfollow(ctx context.Context, arg FindToUnfollowParams) error*/
