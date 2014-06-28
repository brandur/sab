package main

func runShutdown() {
	needApiKey()

	client := newSabClient(*url, *apiKey)
	_, err := client.shutdown()
	if err != nil {
		printFatal("api request: %v", err.Error())
	}
}
