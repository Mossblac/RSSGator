package ext

import (
	"context"
	"fmt"
)

func ScrapeFeeds(s *State) error {
	nextFetch, err := s.DataBase.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = s.DataBase.MarkFeedFetched(context.Background(), nextFetch.ID)
	if err != nil {
		return err
	}

	fetchedfeed, err := FetchFeed(context.Background(), nextFetch.Url)
	if err != nil {
		return err
	}
	maintitle := fetchedfeed.Channel.Title
	fmt.Printf("Channel: %v\n\n", maintitle)

	for i := range fetchedfeed.Channel.Item {
		title := fetchedfeed.Channel.Item[i].Title

		fmt.Printf("article: %v added\n\n", title)
	}
	return nil
}

// func (q *Queries) GetNextFeedToFetch(ctx context.Context) (Feed, error)
//func (q *Queries) MarkFeedFetched(ctx context.Context, id int32) error
//-- in ext already - func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error)
