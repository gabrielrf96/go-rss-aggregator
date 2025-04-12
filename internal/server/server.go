package server

import (
	"log"
	"net/http"

	"github.com/gabrielrf96/go-practice-rss-aggregator/internal/app"
	"github.com/gabrielrf96/go-practice-rss-aggregator/internal/handler"
)

func StartAPIServer(h *handler.Handler, a *app.App) {
	server := &http.Server{
		Handler: getRouter(h, a),
		Addr:    ":" + a.Config.Port,
	}

	log.Printf("[SERVER] Starting on port %s.", a.Config.Port)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
