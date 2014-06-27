package main

import (
	"fmt"
	"strings"

	flag "github.com/ogier/pflag"
)

var (
	numHistories *int
)

func init() {
	numHistories = flag.IntP("num", "n", 10, "number of histories to fetch")
}

func runHistory() {
	needApiKey()

	client := newSabClient(*url, *apiKey)
	histories, err := client.getHistories(*numHistories)
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
