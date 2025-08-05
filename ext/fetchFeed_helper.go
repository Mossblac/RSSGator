package ext

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

/*fetch the rss feed from given url, returns a fill-out
RSSFeed struct*/

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	var rss RSSFeed

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("request with context failed: %v", err)
	}

	req.Header.Set("User-Agent", "rssgator")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("client action failed: %v", err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("read failed: %v", err)
	}

	err = xml.Unmarshal(data, &rss)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("unmarshal error: %v", err)
	}

	rss.Channel.Title = html.UnescapeString(rss.Channel.Title)

	for i := range rss.Channel.Item {
		rss.Channel.Item[i].Title = html.UnescapeString(rss.Channel.Item[i].Title)
		rss.Channel.Item[i].Description = html.UnescapeString(rss.Channel.Item[i].Description)
	}

	return &rss, nil
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
