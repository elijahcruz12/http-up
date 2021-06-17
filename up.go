package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("http://example.com")

	var first string

	_ = resp

	if os.Args != nil && len(os.Args) > 1 {
		first = os.Args[1]
	} else {
		first = "unknown"
	}

	if first == "unknown" {
		getLongAnswer(err)
	} else if first == "long" {
		getLongAnswer(err)
	} else if first == "short" {
		getShortAnswer(err)
	} else {
		fmt.Println("Option not defined")
	}

	getLongAnswer(err)
}

func getLongAnswer(err interface{}) {

	if err != nil {
		fmt.Println("Internet is down")

		os.Exit(1)
	} else {
		fmt.Println("Internet is working")

		os.Exit(0)
	}
}

func getShortAnswer(err interface{}) {
	if err != nil {
		fmt.Println("down")

		os.Exit(1)
	} else {
		fmt.Println("up")

		os.Exit(0)
	}
}
