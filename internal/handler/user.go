package handler

import (
	"fmt"
	"net/http"
	"time"

	v "github.com/RussellLuo/validating/v3"
	"github.com/gabrielrf96/go-rss-aggregator/internal/app"
	"github.com/gabrielrf96/go-rss-aggregator/internal/auth"
	"github.com/gabrielrf96/go-rss-aggregator/internal/database"
	"github.com/gabrielrf96/go-rss-aggregator/internal/request"
	"github.com/gabrielrf96/go-rss-aggregator/internal/response"
	"github.com/google/uuid"
)

type UserResponse struct {
	Name      string     `json:"name"`
	APIKey    string     `json:"secret,omitempty"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type CreateUserParams struct {
	Name string `json:"name"`
}

func (p *CreateUserParams) Schema() v.Schema {
	return v.Schema{
		v.F("name", p.Name): v.Nonzero[string]().Msg(request.ValidationRequired),
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	params, err := getParamsOrRespondWithError[*CreateUserParams](w, r)
	if err != nil {
		return
	}

	apiKey, err := auth.GenerateAPIKey()
	if err != nil {
		response.RespondWithError(w, app.NewInternalError(
			fmt.Sprintf("Could not create API key: %v", err),
		))

		return
	}

	userID := uuid.New()

	user, err := h.a.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        userID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		APIKey:    apiKey.Hash,
	})
	if err != nil {
		response.RespondWithError(w, app.NewInternalError(
			fmt.Sprintf("Could not create user: %v", err),
		))

		return
	}

	response.RespondWithJson(w, http.StatusCreated, UserResponse{
		Name:      user.Name,
		APIKey:    fmt.Sprintf("%s-%s", userID, apiKey.Value),
		CreatedAt: &user.CreatedAt,
	})
}

func (*Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := getUserOrRespondWithError(w, r)
	if err != nil {
		return
	}

	response.RespondWithJson(w, http.StatusOK, UserResponse{
		Name:      user.Name,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
	})
}
