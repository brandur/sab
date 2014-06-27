package main

import (
	"fmt"
)

func runJobs() {
	needApiKey()

	client := newSabClient(*url, *apiKey)
	jobs, err := client.getJobs()
	if err != nil {
		printFatal("api request: %v", err.Error())
	}

	for _, job := range jobs {
		fmt.Printf("%v %v %v %v\n", *job.Filename, job.Size, job.SizeLeft, *job.TimeLeft)
	}
}
