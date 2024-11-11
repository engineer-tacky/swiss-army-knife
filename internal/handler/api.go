package handler

import (
	"net/http"

	"github.com/engineer-tacky/swiss-army-knife/internal/utils"
)

func JSONIndentHandler(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("input")
	json := utils.IndentJSON(input)
	w.Write([]byte(json))
}
