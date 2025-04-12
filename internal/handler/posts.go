package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gabrielrf96/go-practice-rss-aggregator/internal/app"
	"github.com/gabrielrf96/go-practice-rss-aggregator/internal/database"
	"github.com/gabrielrf96/go-practice-rss-aggregator/internal/response"
	"github.com/google/uuid"
)

type PostResponse struct {
	FeedID      uuid.UUID `json:"feed_id"`
	Feed        string    `json:"feed"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	URL         string    `json:"url"`
	PublishedAt time.Time `json:"published_at"`
}

func (h *Handler) GetPosts(w http.ResponseWriter, r *http.Request) {
	user, err := getUserOrRespondWithError(w, r)
	if err != nil {
		return
	}

	posts, err := h.a.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(h.a.Config.API.ReturnPosts),
	})
	if err != nil {
		response.RespondWithError(w, app.NewInternalError(
			fmt.Sprintf("Could not get posts for user \"%s\": %v", user.ID, err),
		))
	}

	responseItems := make([]PostResponse, 0, len(posts))
	for _, post := range posts {
		description := &post.Description.String
		if !post.Description.Valid {
			description = nil
		}

		responseItems = append(responseItems, PostResponse{
			FeedID:      post.FeedID,
			Feed:        post.Feed.Name,
			Title:       post.Title,
			Description: description,
			URL:         post.URL,
			PublishedAt: post.PublishedAt,
		})
	}

	response.RespondWithJson(w, http.StatusOK, responseItems)
}
