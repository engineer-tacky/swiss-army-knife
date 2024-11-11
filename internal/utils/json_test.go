package utils_test

import (
	"testing"

	"github.com/engineer-tacky/swiss-army-knife/internal/utils"
)

func TestIndentJSON(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect string
	}{
		{
			name:   "simple json",
			input:  `{"name": "taro", "age": 20}`,
			expect: "{\n  \"name\": \"taro\",\n  \"age\": 20\n}",
		},
		{
			name:   "nested json",
			input:  `{"name": "taro", "age": 20, "address": {"country": "Japan", "city": "Tokyo"}}`,
			expect: "{\n  \"name\": \"taro\",\n  \"age\": 20,\n  \"address\": {\n    \"country\": \"Japan\",\n    \"city\": \"Tokyo\"\n  }\n}",
		},
		{
			name:   "invalid json",
			input:  `{"name": "taro", "age": 20`,
			expect: "JSON unmarshal error: unexpected end of JSON input",
		},
		{
			name:   "empty json",
			input:  `{}`,
			expect: "{}",
		},
		{
			name:   "empty string",
			input:  ``,
			expect: "JSON unmarshal error: unexpected end of JSON input",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := utils.IndentJSON(tt.input)
			if actual != tt.expect {
				t.Errorf("unexpected result: actual: %v, expect: %v", actual, tt.expect)
			}
		})
	}
}
