package request

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gabrielrf96/go-rss-aggregator/internal/database"
)

type contextKey string

const (
	contextUser contextKey = "user"
)

type ContextError struct {
	contextKey contextKey
}

func (err *ContextError) Error() string {
	return fmt.Sprintf("Failed to retrieve context value for '%s': type assertion failed", err.contextKey)
}

func WithUser(r *http.Request, user *database.User) context.Context {
	return context.WithValue(r.Context(), contextUser, user)
}

func GetUser(r *http.Request) (*database.User, error) {
	user, ok := r.Context().Value(contextUser).(*database.User)
	if !ok {
		return nil, &ContextError{contextKey: contextUser}
	}

	return user, nil
}
