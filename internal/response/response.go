package response

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gabrielrf96/go-rss-aggregator/internal/app"
)

type ErrorResponse struct {
	Error       string   `json:"error"`
	Corrections []string `json:"corrections,omitempty"`
}

func RespondWithJson(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func RespondWithError(w http.ResponseWriter, err *app.APIError) {
	if err.LogMessage != "" {
		log.Printf("[HTTP %d] %s", err.HttpCode, err.LogMessage)
	}

	RespondWithJson(w, err.HttpCode, ErrorResponse{
		Error:       err.Message,
		Corrections: err.Corrections,
	})
}
