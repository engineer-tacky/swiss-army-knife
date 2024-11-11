package main

import (
	"flag"
	"fmt"

	"github.com/engineer-tacky/swiss-army-knife/internal/utils"
)

const (
	JSON   = "json"
	INDENT = "indent"
)

func main() {
	var inputType string
	var action string
	var inputString string

	flag.StringVar(&inputType, "t", "", "Specify the type of the utility to use")
	flag.StringVar(&action, "a", "", "Specify the action to perform")
	flag.StringVar(&inputString, "s", "", "Specify the input string")
	flag.Parse()

	if inputType == JSON && action == INDENT {
		input := inputString
		json := utils.IndentJSON(input)
		fmt.Println(json)
	}
}
