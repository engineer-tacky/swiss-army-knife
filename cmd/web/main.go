package main

import (
	"net/http"

	"github.com/engineer-tacky/swiss-army-knife/internal/handler"
)

func main() {
	// json IndentJSON のルーティング(POST)
	http.HandleFunc("/json/indent", handler.JSONIndentHandler)

	// サーバーを起動
	http.ListenAndServe(":8080", nil)
}
