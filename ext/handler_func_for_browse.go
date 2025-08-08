package ext

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Mossblac/RSSGator/internal/database"
)

func HandlerBrowse(s *State, cmd Command) error {

	var numposts int
	var err error

	if len(cmd.Args) == 0 {
		numposts = 2
	} else {

		numposts, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("unable to convert input to integer: %v", err)
		}
	}
	user, err := s.DataBase.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return err
	}

	userParam := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(numposts),
	}

	posts, err := s.DataBase.GetPostsForUser(context.Background(), userParam)
	if err != nil {
		return err
	}

	for i := range posts {
		fmt.Printf("%v:   %v\n\n", posts[i].FeedName, posts[i].Title)
		fmt.Printf("%v\n\n", posts[i].Url)
		fmt.Printf("%v\n\n\n\n\n", posts[i].Description)

	}

	return nil

}

/*type GetPostsForUserParams struct {
	UserID uuid.UUID
	Limit  int32
}*/

// func (q *Queries) GetPostsForUser(ctx context.Context, arg GetPostsForUserParams) ([]GetPostsForUserRow, error)

/*type GetPostsForUserRow struct {
	ID          int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Url         string
	Description string
	PublishedAt time.Time
	FeedID      int32
	FeedName    string
}*/
