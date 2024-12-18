package utils

import (
	"encoding/json"
	"fmt"
)

func IndentJson(input string) string {
	var rawMessage json.RawMessage

	err := json.Unmarshal([]byte(input), &rawMessage)
	if err != nil {
		return fmt.Sprintf("JSON unmarshal error: %v", err)
	}

	res, err := json.MarshalIndent(&rawMessage, "", "  ")
	if err != nil {
		return fmt.Sprintf("JSON marshal error: %v", err)
	}

	return string(res)
}
