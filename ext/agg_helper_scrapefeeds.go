package ext

import (
	"context"
	"fmt"
	"time"

	"github.com/Mossblac/RSSGator/internal/database"
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
	fmt.Printf("%+v", fetchedfeed)

	for i := range fetchedfeed.Channel.Item {
		title := fetchedfeed.Channel.Item[i].Title
		link := fetchedfeed.Channel.Item[i].Link
		description := fetchedfeed.Channel.Item[i].Description
		pubdate := fetchedfeed.Channel.Item[i].PubDate

		CreatePostParam := database.CreatePostsParams{
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       title,
			Url:         link,
			Description: description,
			PublishedAt: pubdate,
			FeedID:      nextFetch.ID,
		}

		post, err := s.DataBase.CreatePosts(context.Background(), CreatePostParam)
		if err != nil {
			return fmt.Errorf("unable to create post: %v", err)
		}
		fmt.Printf("post entry created: %v\n\n", post.Title)
	}
	return nil
}

/*type CreatePostsParams struct {
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
	Title       string
	Url         string
	Description sql.NullString
	PublishedAt sql.NullTime
	FeedID      int32
}

func (q *Queries) CreatePosts(ctx context.Context, arg CreatePostsParams)*/

//func (q *Queries) GetPostsForUser(ctx context.Context, limit int32) ([]Post, error)
