package server

import (
	"encoding/xml"
	"io"

	"github.com/gabrielrf96/go-practice-rss-aggregator/internal/app"
)

type RSSFeed struct {
	Channel RSSChannel `xml:"channel"`
}

type RSSChannel struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Language    string    `xml:"language"`
	Item        []RSSItem `xml:"item"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urlToFeed(url string, a *app.App) (RSSFeed, error) {
	rssFeed := RSSFeed{}

	response, err := a.HTTPClient.Get(url)
	if err != nil {
		return rssFeed, err
	}
	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return rssFeed, err
	}

	err = xml.Unmarshal(bytes, &rssFeed)
	if err != nil {
		return rssFeed, err
	}

	return rssFeed, nil
}
