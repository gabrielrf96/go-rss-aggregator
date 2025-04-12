package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gabrielrf96/go-rss-aggregator/internal/app"
	"github.com/gabrielrf96/go-rss-aggregator/internal/database"
	"github.com/gabrielrf96/go-rss-aggregator/internal/request"
	"github.com/gabrielrf96/go-rss-aggregator/internal/response"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type SubscriptionResponse struct {
	FeedID    uuid.UUID  `json:"feed_id"`
	Name      string     `json:"name,omitempty"`
	URL       string     `json:"url,omitempty"`
	UserID    *uuid.UUID `json:"user_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (h *Handler) GetSubscriptions(w http.ResponseWriter, r *http.Request) {
	user, err := getUserOrRespondWithError(w, r)
	if err != nil {
		return
	}

	subscriptions, err := h.a.DB.GetActiveSubscriptions(r.Context(), user.ID)
	if err != nil {
		response.RespondWithError(w, app.NewInternalError(
			fmt.Sprintf("Could not retrieve subscriptions: %v", err),
		))

		return
	}

	responseItems := make([]SubscriptionResponse, 0, len(subscriptions))
	for _, subscription := range subscriptions {
		responseItems = append(responseItems, SubscriptionResponse{
			FeedID:    subscription.FeedID,
			Name:      subscription.Feed.Name,
			URL:       subscription.Feed.URL,
			CreatedAt: &subscription.CreatedAt,
		})
	}

	response.RespondWithJson(w, http.StatusOK, responseItems)
}

func (h *Handler) Subscribe(w http.ResponseWriter, r *http.Request) {
	feedID, err := getFeedIdOrRespondWithError(w, r)
	if err != nil {
		return
	}

	user, err := getUserOrRespondWithError(w, r)
	if err != nil {
		return
	}

	feed, err := h.a.DB.GetFeed(r.Context(), feedID)
	if err != nil {
		if err == sql.ErrNoRows {
			response.RespondWithError(w, app.NewNotFoundError("The provided feed does not exist"))
		} else {
			response.RespondWithError(w, newGenericSubscriptionError(err))
		}

		return
	}

	subscription, err := h.a.DB.CreateSubscription(r.Context(), database.CreateSubscriptionParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    feedID,
		UserID:    user.ID,
	})
	if err != nil {
		if database.IsError(err, database.UniqueViolation) {
			response.RespondWithError(w, app.NewConflictError("You are already subscribed to that feed"))
		} else {
			response.RespondWithError(w, newGenericSubscriptionError(err))
		}

		return
	}

	response.RespondWithJson(w, http.StatusCreated, SubscriptionResponse{
		FeedID:    subscription.FeedID,
		Name:      feed.Name,
		URL:       feed.URL,
		CreatedAt: &subscription.CreatedAt,
	})
}

func (h *Handler) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	feedID, err := getFeedIdOrRespondWithError(w, r)
	if err != nil {
		return
	}

	user, err := getUserOrRespondWithError(w, r)
	if err != nil {
		return
	}

	now := time.Now().UTC()
	count, err := h.a.DB.DeleteSubscription(r.Context(), database.DeleteSubscriptionParams{
		UserID: user.ID,
		FeedID: feedID,
	})
	if err != nil {
		response.RespondWithError(w, app.NewInternalError(
			fmt.Sprintf("Could not unsubscribe: %v", err),
		))

		return
	}

	if count == 0 {
		response.RespondWithError(w, app.NewNotFoundError("You are not subscribed to the provided feed"))
	} else {
		response.RespondWithJson(w, http.StatusOK, SubscriptionResponse{
			FeedID:    feedID,
			DeletedAt: &now,
		})
	}
}

func getFeedIdOrRespondWithError(w http.ResponseWriter, r *http.Request) (uuid.UUID, error) {
	feedID, err := uuid.Parse(chi.URLParam(r, request.URLParamFeedID))
	if err != nil {
		response.RespondWithError(w, &app.APIError{
			Message:  "Failed parsing feed ID, make sure you provided a valid ID",
			HttpCode: http.StatusBadRequest,
		})

		return uuid.Nil, err
	}

	return feedID, nil
}

func newGenericSubscriptionError(err error) *app.APIError {
	return app.NewInternalError(
		fmt.Sprintf("Could not subscribe to the feed: %v", err),
	)
}
