package ext

import (
	"context"
)

/*fetch the rss feed from given url, returns a fill-out
RSSFeed struct*/

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	return &RSSFeed{}, nil
}

/*type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

*/

/*type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}
*/
