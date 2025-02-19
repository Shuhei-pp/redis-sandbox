package usecase

import (
	"api/internal/lib"
	"context"
	"fmt"
	"net/http"
)

func GetRedisValue(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	key := r.URL.Query().Get("key")

	var client = lib.RedisClient()

	ctx := context.Background()

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		http.Error(w, "Redis GET error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("GET %s: %s", key, val)))
}
