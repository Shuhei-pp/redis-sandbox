package usecase

import (
	"api/internal/lib"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type Request struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func SetRedisValue(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var req Request
	if err = json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Failed to unmarshal request body", http.StatusInternalServerError)
		return
	}

	redisClient := lib.RedisClient()
	ctx := context.Background()

	if _, err = redisClient.Set(ctx,
		req.Key,
		req.Value,
		0).Result(); err != nil {
		http.Error(w, "Redis SET error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("SET key: value"))
}
