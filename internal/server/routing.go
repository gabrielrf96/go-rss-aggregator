package server

import (
	"fmt"
	"net/http"

	"github.com/gabrielrf96/go-practice-rss-aggregator/internal/app"
	"github.com/gabrielrf96/go-practice-rss-aggregator/internal/auth"
	"github.com/gabrielrf96/go-practice-rss-aggregator/internal/handler"
	"github.com/gabrielrf96/go-practice-rss-aggregator/internal/request"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

var corsOptions = cors.Options{
	AllowedOrigins: []string{"https://*", "http://*"},
	AllowedMethods: []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodOptions,
	},
	AllowedHeaders:   []string{"*"},
	ExposedHeaders:   []string{"Link"},
	AllowCredentials: false,
	MaxAge:           300,
}

func getRouter(h *handler.Handler, a *app.App) chi.Router {
	authMiddleware := auth.NewAuthMiddleware(a)

	r := chi.NewRouter()

	r.Use(cors.Handler(corsOptions))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/healthz", h.Healthz)

		r.Route("/user", func(r chi.Router) {
			r.Post("/", h.CreateUser)

			// Authenticated endpoints
			r.Group(func(r chi.Router) {
				r.Use(authMiddleware)
				r.Get("/", h.GetUser)
			})
		})

		r.Route("/feeds", func(r chi.Router) {
			r.Use(authMiddleware)

			r.Post("/", h.CreateFeed)
			r.Get("/", h.GetFeeds)
		})

		r.Route("/subscriptions", func(r chi.Router) {
			r.Use(authMiddleware)

			parameterizedFeedURL := fmt.Sprintf("/{%s}", request.URLParamFeedID)

			r.Get("/", h.GetSubscriptions)
			r.Post(parameterizedFeedURL, h.Subscribe)
			r.Delete(parameterizedFeedURL, h.Unsubscribe)
		})

		r.Route("/posts", func(r chi.Router) {
			r.Use(authMiddleware)

			r.Get("/", h.GetPosts)
		})
	})

	return r
}
