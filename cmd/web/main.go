package main

import (
	"net/http"

	h "github.com/engineer-tacky/swiss-army-knife/internal/handler"
	m "github.com/engineer-tacky/swiss-army-knife/internal/middleware"
)

func main() {
	http.HandleFunc("/json/indent", m.Cors(h.JsonIndentHandler))

	http.ListenAndServe(":8080", nil)
}
