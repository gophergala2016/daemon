package reader

import "testing"

var feeds = []string{
	"http://feeds.reuters.com/reuters/topnews?format=xml",
	"http://www.npr.org/rss/rss.php?id=1001",
	"http://rss.cnn.com/rss/cnn_latest.rss",
	"http://rss.ireport.com/feeds/oncnn.rss",
	"http://rss.nytimes.com/services/xml/rss/nyt/World.xml",
	"http://feeds.bbci.co.uk/news/rss.xml",
	"http://www.un.org/apps/news/rss/rss_top.asp",
}

func TestGetFeed(t *testing.T) {
	// TODO: Fix where it looks for the `charsets.json` file
	// for _, url := range feeds {
	// 	feed, err := GetFeed(url)
	// 	if err != nil {
	// 		t.Errorf("(%s) %s", url, err.Error())
	// 	} else if feed.Title == "" {
	// 		t.Errorf("Feed (%s) has no title", url)
	// 	}
	// }
}
