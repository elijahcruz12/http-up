package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Testing your internet connection")

	resp, err := http.Get("http://example.com")

	_ = resp

	if err != nil {
		fmt.Println("Internet is down")
	} else {
		fmt.Println("Internet is working")

	}
}
