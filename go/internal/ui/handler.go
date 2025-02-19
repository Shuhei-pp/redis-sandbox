package ui

import (
	"api/internal/usecase"
	"net/http"
)

func Handler() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello,go"))
	})

	http.HandleFunc("/redis/get", usecase.GetRedisValue)
	http.HandleFunc("/redis/set", usecase.SetRedisValue)
	http.HandleFunc("/product/list", usecase.GetProductList)

	http.ListenAndServe(":8080", nil)
}
