package reader

import (
	"io/ioutil"
	"net/http"
)

// Scrape the feed contents from the given URI
func scrape(url string) ([]byte, error) {
	// Retrieve the feed
	response, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()

	// Read the contents
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}
	return contents, nil
}

// Get the feed for an RSS endpoint
func GetFeed(url string) (Rss, error) {
	content, err := scrape(url)
	if err != nil {
		return Rss{}, err
	}

	feed := Rss{}
	err = feed.ParseFeed(content)
	return feed, err
}
