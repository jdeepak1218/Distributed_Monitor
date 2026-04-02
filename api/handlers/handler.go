package handlers

import (
	"encoding/json"
	"net/http"

	goredis "github.com/go-redis/redis/v8"
)

type URLRequest struct {
	URL string `json:"url"`
}
type Handler struct {
	redisClient *goredis.Client
}

func NewHandler(client *goredis.Client) *Handler {
	return &Handler{redisClient: client}
}
func (h *Handler) HealthCheckup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
