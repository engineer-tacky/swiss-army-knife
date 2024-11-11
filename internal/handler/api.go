package handler

import (
	"net/http"

	"github.com/engineer-tacky/swiss-army-knife/internal/utils"
)

func JSONIndentHandler(w http.ResponseWriter, r *http.Request) {
	// POSTされたJSONを取得
	input := r.FormValue("input")
	// JSONのインデントを整形
	json := utils.IndentJSON(input)
	// 整形したJSONをレスポンスとして返す
	w.Write([]byte(json))
}
