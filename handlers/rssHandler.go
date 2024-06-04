package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"rssProxy/constants"
	"rssProxy/rssCache"

	"github.com/SlyMarbo/rss"
	"github.com/patrickmn/go-cache"
)

type Feed struct {
	Nickname    string   `json:"nickname"`
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Description string   `json:"description"`
	Link        string   `json:"link"`
	UpdateURL   string   `json:"updatedURL"`
	Categories  []string `json:"categories"`
	Items       []Item   `json:"items"`
}

type Item struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Link    string `json:"link"`
}

func GetRss(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(constants.HEADER_KEY_CONTENT_TYPE, constants.HEADER_VALUE_APPLICATION_JSON)
	w.Header().Add(constants.HEADER_CORS_ACCESS_CONTROL, "*")

	var rssUrl = r.URL.Query().Get("url")

	if len(rssUrl) == 0 {
		log.Default().Println("No url argument passed in query parameter")
		http.Error(w, "No URL in Query Parameter", http.StatusBadRequest)
	}

	log.Default().Printf("Fetching RSS feed for %s\n", rssUrl)

	var cacheResult, found = rssCache.GetCache().Get(rssUrl)

	if found {
		log.Default().Println("sending from cache")
		var response, convertok = cacheResult.(Feed)
		if !convertok {
			log.Default().Printf("Could not convert cache entry to type rss.Feed for %s\n", rssUrl)
		}
		json.NewEncoder(w).Encode(response)
	} else {
		var feed, err = rss.Fetch(rssUrl)
		errHandler(err, w)
		var feedEntry = Feed{Nickname: feed.Nickname,
			Title:       feed.Title,
			Author:      feed.Author,
			Description: feed.Description,
			Link:        feed.Link,
			UpdateURL:   feed.UpdateURL,
			Categories:  feed.Categories,
		}

		var items = []Item{}
		for _, item := range feed.Items {
			items = append(items, Item{Title: item.Title, Summary: item.Summary, Link: item.Link})
		}

		feedEntry.Items = items

		rssCache.GetCache().Set(rssUrl, feedEntry, cache.DefaultExpiration)
		json.NewEncoder(w).Encode(feedEntry)
	}

}

func errHandler(err error, w http.ResponseWriter) {
	if err != nil {
		log.Default().Println(err)
		http.Error(w, err.Error(), 500)
	}
}
