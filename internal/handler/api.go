package handler

import (
	"net/http"

	"github.com/engineer-tacky/swiss-army-knife/internal/utils"
)

func JsonIndentHandler(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("input")
	json := utils.IndentJson(input)
	w.Write([]byte(json))
}
