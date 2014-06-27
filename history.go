package main

import (
	"fmt"
	"strings"
)

func runHistory() {
	needApiKey()

	client := newSabClient(*url, *apiKey)
	histories, err := client.getHistories(10)
	if err != nil {
		printFatal("api request: %v", err.Error())
	}

	for _, history := range histories {
		listRec(strings.ToLower(*history.Status), prettySize(history.Size), *history.Name)
	}
}

func prettySize(size int) string {
	return fmt.Sprintf("%vM", size/1024/1024)
}
