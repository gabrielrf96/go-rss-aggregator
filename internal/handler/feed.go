package handler

import (
	"fmt"
	"net/http"
	"time"

	v "github.com/RussellLuo/validating/v3"
	"github.com/RussellLuo/vext"
	"github.com/gabrielrf96/go-practice-rss-aggregator/internal/app"
	"github.com/gabrielrf96/go-practice-rss-aggregator/internal/database"
	"github.com/gabrielrf96/go-practice-rss-aggregator/internal/request"
	"github.com/gabrielrf96/go-practice-rss-aggregator/internal/response"
	"github.com/google/uuid"
)

type FeedResponse struct {
	ID        uuid.UUID  `json:"id,omitempty"`
	Name      string     `json:"name"`
	URL       string     `json:"url"`
	CreatedAt *time.Time `json:"created_at"`
	FetchedAt *time.Time `json:"fetched_at"`
}

type CreateFeedParams struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (p *CreateFeedParams) Schema() v.Schema {
	return v.Schema{
		v.F("name", p.Name): v.Nonzero[string]().Msg(request.ValidationRequired),
		v.F("url", p.URL): v.All(
			v.Nonzero[string]().Msg(request.ValidationRequired),
			vext.URL().Msg(request.ValidationURL),
		),
	}
}

func (h *Handler) CreateFeed(w http.ResponseWriter, r *http.Request) {
	params, err := getParamsOrRespondWithError[*CreateFeedParams](w, r)
	if err != nil {
		return
	}

	user, err := getUserOrRespondWithError(w, r)
	if err != nil {
		return
	}

	feed, err := h.a.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		URL:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		if database.IsError(err, database.UniqueViolation) {
			response.RespondWithError(w, app.NewConflictError("The provided feed already exists"))
		} else {
			response.RespondWithError(w, app.NewInternalError(
				fmt.Sprintf("Could not create feed: %v", err),
			))
		}

		return
	}

	response.RespondWithJson(w, http.StatusCreated, FeedResponse{
		ID:        feed.ID,
		Name:      feed.Name,
		URL:       feed.URL,
		CreatedAt: &feed.CreatedAt,
	})
}

func (h *Handler) GetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := h.a.DB.GetFeeds(r.Context())
	if err != nil {
		response.RespondWithError(w, app.NewInternalError(err.Error()))
	}

	responseItems := make([]FeedResponse, 0, len(feeds))
	for _, feed := range feeds {
		fetchedAt := &feed.FetchedAt.Time

		if !feed.FetchedAt.Valid {
			fetchedAt = nil
		}

		responseItems = append(responseItems, FeedResponse{
			ID:        feed.ID,
			Name:      feed.Name,
			URL:       feed.URL,
			CreatedAt: &feed.CreatedAt,
			FetchedAt: fetchedAt,
		})
	}

	response.RespondWithJson(w, http.StatusOK, responseItems)
}
