package app

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type ScraperConfig struct {
	// Max time, in seconds, to wait for a response for the scraper HTTP client.
	Timeout time.Duration
	// Number of concurrent requests the scraper is allowed to make.
	Instances int
	// How often, in seconds, the scraper should fetch feeds.
	Interval time.Duration
}

type APIConfig struct {
	// Number of latest posts to return for the posts endpoint.
	ReturnPosts int
}

type Config struct {
	Port      string
	DBConnUrl string
	Scraper   ScraperConfig
	API       APIConfig
}

func loadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file. Current session environment will be used instead.")
	}

	return &Config{
		Port:      getEnv("PORT"),
		DBConnUrl: getEnv("DB_URL"),
		Scraper: ScraperConfig{
			Timeout:   time.Duration(getEnvInt("SCRAPER_TIMEOUT_SECONDS")) * time.Second,
			Instances: getEnvInt("SCRAPER_INSTANCES"),
			Interval:  time.Duration(getEnvInt("SCRAPER_INTERVAL_SECONDS")) * time.Second,
		},
		API: APIConfig{
			ReturnPosts: getEnvInt("API_RETURN_POSTS"),
		},
	}
}

func getEnv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalf("%s is not defined in the environment.", name)
	}

	return value
}

func getEnvInt(name string) int {
	valueStr := getEnv(name)
	if valueStr == "" {
		log.Fatalf("%s is not defined in the environment.", name)
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Fatalf("%s must be an integer.", name)
	}

	return value
}
