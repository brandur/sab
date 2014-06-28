package main

func runPause() {
	needApiKey()

	client := newSabClient(*url, *apiKey)
	_, err := client.pause()
	if err != nil {
		printFatal("api request: %v", err.Error())
	}
}
