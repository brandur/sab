package main

import (
	"fmt"
)

func runStatus() {
	needApiKey()

	client := newSabClient(url, apiKey)
	status, err := client.getStatus()
	if err != nil {
		printFatal("api request: %v", err.Error())
	}

	fmt.Printf("State: %v\n", *status.State)
}
