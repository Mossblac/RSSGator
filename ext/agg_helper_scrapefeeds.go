package ext

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Mossblac/RSSGator/internal/database"
)

func ScrapeFeeds(s *State) error {
	nextFetch, err := s.DataBase.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("obtained next fetch")

	err = s.DataBase.MarkFeedFetched(context.Background(), nextFetch.ID)
	if err != nil {
		return err
	}
	fmt.Println("marked as fetched")

	fetchedfeed, err := FetchFeed(context.Background(), nextFetch.Url)
	if err != nil {
		return err
	}
	fmt.Println("feed fetched")

	if len(fetchedfeed.Channel.Item) > 0 {

		for i := range fetchedfeed.Channel.Item {
			title := fetchedfeed.Channel.Item[i].Title
			link := fetchedfeed.Channel.Item[i].Link
			description := fetchedfeed.Channel.Item[i].Description
			pubdate := fetchedfeed.Channel.Item[i].PubDate

			pubTime, err := parseTime(pubdate)
			if err != nil {
				log.Printf("Warning: %v, using current time as published time", err)
			}

			CreatePostParam := database.CreatePostsParams{
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				Title:       title,
				Url:         link,
				Description: description,
				PublishedAt: pubTime,
				FeedID:      nextFetch.ID,
			}

			fmt.Printf("Creating post: %+v\n", CreatePostParam)
			fmt.Printf("created at, timestamp format: %v\n\n", CreatePostParam.CreatedAt)
			fmt.Printf("PublishedAt, time format: %v\n\n", CreatePostParam.PublishedAt)

			post, err := s.DataBase.CreatePosts(context.Background(), CreatePostParam)
			if err != nil {
				fmt.Printf("create posts error: %+v\n", err)
				return fmt.Errorf("unable to create post: %v", err)
			}
			fmt.Printf("post entry created: %v\n\n", post.Title)
		}

	} else {
		return fmt.Errorf("items empty")
	}
	return nil
}

func parseTime(publishedAt string) (time.Time, error) {
	t, err := time.Parse(time.RFC1123Z, publishedAt)
	if err != nil {
		return time.Now(), fmt.Errorf("time parsing failed: %v", err)
	}
	return t, nil

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
