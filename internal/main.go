package rssagg

import (
	"github.com/gabrielrf96/go-practice-rss-aggregator/internal/app"
	"github.com/gabrielrf96/go-practice-rss-aggregator/internal/handler"
	"github.com/gabrielrf96/go-practice-rss-aggregator/internal/server"
	_ "github.com/lib/pq"
)

func StartApp() {
	a := app.NewApp()

	go server.StartScraping(a)
	server.StartAPIServer(handler.NewHandler(a), a)
}
