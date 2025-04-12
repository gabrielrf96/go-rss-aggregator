package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gabrielrf96/go-practice-rss-aggregator/internal/database"
)

type Middleware func(http.Handler) http.Handler

// Container for common dependencies and services
type App struct {
	Config     *Config
	DB         *database.Queries
	HTTPClient *http.Client
}

func NewApp() *App {
	config := loadConfig()

	conn, err := sql.Open("postgres", config.DBConnUrl)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	return &App{
		Config: config,
		DB:     database.New(conn),
		HTTPClient: &http.Client{
			Timeout: config.Scraper.Timeout,
		},
	}
}
