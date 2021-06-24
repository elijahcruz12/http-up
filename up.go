package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get("http://example.com")

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
	} else if first == "help" {
		helpScreen()
	} else if first == "--help" {
		helpScreen()
	} else {
		fmt.Println("Option not defined")
	}
}

func helpScreen() {
	fmt.Println("Up (http-up) version 1.1.0")
	fmt.Println("By Elijah Cruz")
	fmt.Println()
	fmt.Println("Allows you to check the status of your internet")
	fmt.Println()
	fmt.Printf("")
	fmt.Println("Options:")
	fmt.Println(" - <none>: Equivalent of long")
	fmt.Println(" - long: Displays a longer more human readable version of your internet status.")
	fmt.Println(" - short: Displays either \"up\" or \"down\" depending on your internet is up or down.")
	fmt.Println(" - help: This help message")
}

func getLongAnswer(err interface{}) {

	fmt.Println("Getting you internet status")

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
