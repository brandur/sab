package main

import (
	"fmt"
	"strconv"
	"strings"
)

func runJobs() {
	needApiKey()

	client := newSabClient(*url, *apiKey)
	jobs, err := client.getJobs()
	if err != nil {
		printFatal("api request: %v", err.Error())
	}

	for _, job := range jobs {
		listRec(strings.ToLower(*job.Status), formatSize(job), *job.Percentage+"%", *job.Filename)
	}
}

func formatSize(j *job) string {
	size, err := strconv.ParseFloat(*j.Size, 64)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%vM", int(size))
}
