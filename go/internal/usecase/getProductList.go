package usecase

import (
	"api/internal/lib"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetProductList(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var client = lib.RedisClient()
	ctx := context.Background()

	cacheKey := "product_list"
	cacheData, err := client.Get(ctx, cacheKey).Result()
	if err == nil {
		fmt.Println("Cache Hit")
		fmt.Println("Elapsed Time: ", time.Since(startTime))
		w.Write([]byte(fmt.Sprintf("GET %s: %s", cacheKey, cacheData)))
		return
	}

	res, err := http.Get("google.com")
	if err != nil {
		http.Error(w, "Failed to fetch product list", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	_, err = client.Set(ctx, cacheKey, body, 30*time.Second).Result()
	if err != nil {
		http.Error(w, "Redis GET error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Cache Miss")
	fmt.Println("Elapsed Time: ", time.Since(startTime))

	w.Write([]byte(fmt.Sprintf("GET %s: %s", cacheKey, body)))
}
