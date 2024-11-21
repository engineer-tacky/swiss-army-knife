package main

import (
	"net/http"

	"github.com/engineer-tacky/swiss-army-knife/internal/handler"
)

func main() {
	http.HandleFunc("/json/indent", handler.JsonIndentHandler)

	http.ListenAndServe(":8080", nil)
}
