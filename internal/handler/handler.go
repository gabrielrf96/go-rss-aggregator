package handler

import (
	"errors"
	"net/http"

	"github.com/gabrielrf96/go-practice-rss-aggregator/internal/app"
	"github.com/gabrielrf96/go-practice-rss-aggregator/internal/database"
	"github.com/gabrielrf96/go-practice-rss-aggregator/internal/request"
	"github.com/gabrielrf96/go-practice-rss-aggregator/internal/response"
)

type Handler struct {
	a *app.App
}

func NewHandler(a *app.App) *Handler {
	return &Handler{
		a: a,
	}
}

// Helpers for common actions in handlers, mainly to avoid repeating boilerplate error / response handling

// Uses [request.ParseParams] to extract a params struct from the provided request and, if the request
// also defines a non-empty set of validation rules ([validating.Schema]), then validates the params struct
// using [request.ValidateParams].
//
// If an error happens at any point in this process, it is automatically handled by writing an appropriate
// JSON error response to the provided [http.ResponseWriter]. Errors are also returned so the calling handler
// can abort execution if necessary.
//
// If you need to handle errors or validation in a different way, or avoid sending them to the client,
// use [request.ParseParams] directly instead.
func getParamsOrRespondWithError[T request.RequestParams](w http.ResponseWriter, r *http.Request) (T, error) {
	var apiErr *app.APIError

	params, err := request.ParseParams[T](r)
	if errors.As(err, &apiErr) {
		response.RespondWithError(w, apiErr)

		return params, apiErr
	}

	err = request.ValidateParams(params)
	if errors.As(err, &apiErr) {
		response.RespondWithError(w, apiErr)

		return params, apiErr
	}

	return params, nil
}

// Retrieves the authenticated user from the provided request's context, and automatically handles errors
// by writing an appropriate JSON error response to the provided [http.ResponseWriter]. The error is also returned
// so the calling handler can abort execution if necessary.
func getUserOrRespondWithError(w http.ResponseWriter, r *http.Request) (*database.User, error) {
	user, err := request.GetUser(r)

	if err != nil {
		response.RespondWithError(w, app.NewInternalError(err.Error()))

		return nil, err
	}

	return user, nil
}
