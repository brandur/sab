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
		listRec(*job.Filename, formatSize(job.Size), formatSize(job.SizeLeft), *job.TimeLeft)
	}
}

func formatSize(size float64) string {
	return fmt.Sprintf("%vM", int(size))
}
