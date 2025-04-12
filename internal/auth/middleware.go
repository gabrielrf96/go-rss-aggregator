package auth

import (
	"errors"
	"net/http"

	"github.com/gabrielrf96/go-rss-aggregator/internal/app"
	"github.com/gabrielrf96/go-rss-aggregator/internal/request"
	"github.com/gabrielrf96/go-rss-aggregator/internal/response"
	"golang.org/x/crypto/bcrypt"
)

func NewAuthMiddleware(a *app.App) app.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authData, err := getAuthData(r.Header)
			var apiErr *app.APIError
			if errors.As(err, &apiErr) {
				response.RespondWithError(w, apiErr)

				return
			}

			user, err := a.DB.GetUser(r.Context(), authData.UserID)
			if err != nil {
				respondWithUnauthorizedError(w)

				return
			}

			err = bcrypt.CompareHashAndPassword([]byte(user.APIKey), []byte(authData.APIKey))
			if err != nil {
				respondWithUnauthorizedError(w)

				return
			}

			ctx := request.WithUser(r, &user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func respondWithUnauthorizedError(w http.ResponseWriter) {
	response.RespondWithError(w, &app.APIError{
		HttpCode: http.StatusUnauthorized,
		Message:  "Auth error: unrecognized credentials",
	})
}
