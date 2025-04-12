package handler

import (
	"net/http"

	"github.com/gabrielrf96/go-rss-aggregator/internal/response"
)

type HealthzResponse struct {
	Status string `json:"status"`
}

func (*Handler) Healthz(w http.ResponseWriter, r *http.Request) {
	response.RespondWithJson(w, http.StatusOK, HealthzResponse{Status: "ok"})
}
