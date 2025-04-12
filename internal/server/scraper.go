package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gabrielrf96/go-rss-aggregator/internal/app"
	"github.com/gabrielrf96/go-rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func startScraping(a *app.App) {
	log.Printf(
		"[SCRAPER] Starting scraper for: %v feeds every %v seconds",
		a.Config.Scraper.Instances,
		int(a.Config.Scraper.Interval.Seconds()),
	)

	ticker := time.NewTicker(a.Config.Scraper.Interval)

	for ; ; <-ticker.C {
		feeds, err := a.DB.GetNextFeedsToFetch(context.Background(), int32(a.Config.Scraper.Instances))
		if err != nil {
			log.Printf("[SCRAPER] Error fetching feeds: %v", err)

			continue
		}

		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(a, wg, feed)
		}

		wg.Wait()
	}
}

func scrapeFeed(a *app.App, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := a.DB.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		logErrorFetchingFeed(feed, err)

		return
	}

	rssFeed, err := urlToFeed(feed.URL, a)
	if err != nil {
		logErrorFetchingFeed(feed, err)

		return
	}

	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			logErrorFetchingFeed(feed, fmt.Errorf("failed parsing date \"%s\": %w", item.PubDate, err))

			continue
		}

		_, err = a.DB.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			PublishedAt: pubDate,
			Title:       item.Title,
			Description: description,
			URL:         item.Link,
			FeedID:      feed.ID,
		})
		if err != nil && !database.IsError(err, database.UniqueViolation) {
			logErrorFetchingFeed(feed, fmt.Errorf("failed storing post in DB: %w", err))
		}
	}

	log.Printf("[SCRAPER] Fetched \"%s\", %v posts found", feed.Name, len(rssFeed.Channel.Item))
}

func logErrorFetchingFeed(feed database.Feed, err error) {
	log.Printf("[SCRAPER] Error fetching feed \"%s\": %v", feed.Name, err)
}
