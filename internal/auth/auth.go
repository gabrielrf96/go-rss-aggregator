package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"

	"github.com/gabrielrf96/go-rss-aggregator/internal/app"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type GeneratedAPIKey struct {
	Value string
	Hash  string
}

type AuthData struct {
	UserID uuid.UUID
	APIKey string
}

func GenerateAPIKey() (*GeneratedAPIKey, error) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	tmpHash := sha256.New()
	tmpHash.Write(randomBytes)

	plainAPIKey := hex.EncodeToString(tmpHash.Sum(nil))

	hashedAPIKey, err := hash(plainAPIKey)
	if err != nil {
		return nil, err
	}

	apiKey := &GeneratedAPIKey{
		Value: plainAPIKey,
		Hash:  hashedAPIKey,
	}

	return apiKey, nil
}

// Returns an API key from the provided HTTP headers.
// Expects an Authorization header in the form of "Bearer [API_KEY]".
func getAuthData(headers http.Header) (*AuthData, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return nil, &app.APIError{
			Message:  "Auth error: credentials not provided for authenticated endpoint",
			HttpCode: http.StatusUnauthorized,
		}
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 {
		return nil, &app.APIError{
			Message:  "Auth error: malformed Authorization header",
			HttpCode: http.StatusBadRequest,
		}
	}

	if parts[0] != "Bearer" {
		return nil, &app.APIError{
			Message: "Auth error: malformed Authorization header",
			Corrections: []string{
				fmt.Sprintf("Expected authentication scheme 'Bearer', received '%s'", parts[0]),
			},
			HttpCode: http.StatusBadRequest,
		}
	}

	authDataStr := parts[1]
	separatorIdx := strings.LastIndex(authDataStr, "-")
	if separatorIdx == -1 {
		return nil, newInvalidAPIKeyError()
	}

	userID, err := uuid.Parse(authDataStr[:separatorIdx])
	if err != nil {
		return nil, newInvalidAPIKeyError()
	}

	authData := &AuthData{
		UserID: userID,
		APIKey: authDataStr[separatorIdx+1:],
	}

	return authData, nil
}

func hash(apiKey string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(apiKey), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}

func newInvalidAPIKeyError() *app.APIError {
	return &app.APIError{
		Message:  "Auth error: invalid API key",
		HttpCode: http.StatusUnauthorized,
	}
}
