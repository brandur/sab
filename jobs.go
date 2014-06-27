package main

import (
	"fmt"
)

func jobs() {
	needApiKey()

	client := newSabClient(url, apiKey)
	thing, err := client.getJobs()
	if err != nil {
		printFatal("api request: %v", err.Error())
	}

	fmt.Printf("data: %+v\n", thing)
}
