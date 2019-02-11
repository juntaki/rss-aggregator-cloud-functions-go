package p

import (
	"github.com/rs/cors"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/feeds"
	"github.com/mmcdole/gofeed"
)

func handler(w http.ResponseWriter, r *http.Request) {
	cors.Default().Handler(http.HandlerFunc(feedHandler)).ServeHTTP(w, r)
}

func feedHandler(w http.ResponseWriter, r *http.Request) {

	fp := gofeed.NewParser()
	input1, _ := fp.ParseURL("https://medium.com/feed/@juntaki")
	input2, _ := fp.ParseURL("https://qiita.com/juntaki/feed.atom")

	now := time.Now()
	output := &feeds.Feed{
		Title:       "juntaki.com",
		Link:        &feeds.Link{Href: "https://juntaki.com"},
		Description: "Aggregated feed by juntaki",
		Author:      &feeds.Author{Name: "Jumpei Takiysu", Email: "me@juntaki.com"},
		Created:     now,
	}

	output.Items = []*feeds.Item{}

	for _, item := range append(input1.Items, input2.Items...) {
		output.Items = append(output.Items, &feeds.Item{
			Title:       item.Title,
			Link:        &feeds.Link{Href: item.Link},
			Description: item.Description,
			Author:      &feeds.Author{Name: item.Author.Name, Email: item.Author.Email},
			Created:     *item.PublishedParsed,
		})
	}

	rss, err := output.ToRss()
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(rss))
}
