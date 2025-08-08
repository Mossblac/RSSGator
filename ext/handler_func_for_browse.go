package ext

import (
	"context"
	"fmt"
)

func HandlerBrowse(s *State, cmd Command) error {

	posts, err := s.DataBase.GetPostsForUser(context.Background(), 2)
	if err != nil {
		return err
	}

	for i := range posts {
		feed, err := FetchFeed(context.Background(), posts[i].Url)
		if err != nil {
			return err
		}

		fmt.Printf("%+v", feed.Channel.Item)
	}

	return nil

}

//func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error)
//func (q *Queries) GetPostsForUser(ctx context.Context, limit int32) ([]Post, error)
