package main

func runResume() {
	needApiKey()

	client := newSabClient(*url, *apiKey)
	_, err := client.resume()
	if err != nil {
		printFatal("api request: %v", err.Error())
	}
}
