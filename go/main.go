package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func main() {
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // コンテナ名で接続
		Password: "",
		DB:       0,
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello,go"))
	})

	http.HandleFunc("/redis", func(w http.ResponseWriter, r *http.Request) {

		ctx := context.Background()
		if err := client.Ping(ctx).Err(); err != nil {
			http.Error(w, "Redis connection error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		err := client.Set(ctx, "foo", "bar", 0).Err()

		if err != nil {
			w.Write([]byte(err.Error()))
		}

		val, err := client.Get(ctx, "foo").Result()
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		fmt.Fprintf(w, "foo: %s", val)
	})

	http.ListenAndServe(":8080", nil)
}
