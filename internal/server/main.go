package server

import (
	"github.com/gabrielrf96/go-rss-aggregator/internal/app"
	"github.com/gabrielrf96/go-rss-aggregator/internal/handler"
)

func StartApp() {
	a := app.NewApp()

	go startScraping(a)
	startAPIServer(handler.NewHandler(a), a)
}
