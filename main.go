package main

import (
	"fmt"
	"os"

	flag "github.com/ogier/pflag"
)

var (
	apiKey  *string
	url     *string
	verbose *bool
)

func main() {
	parseFlags()

	if len(flag.Args()) != 1 {
		flag.Usage()
	}

	switch flag.Args()[0] {
	case "history":
		runHistory()
	case "pause":
		runPause()
	case "queue":
		runQueue()
	case "resume":
		runResume()
	case "shutdown":
		runShutdown()
	case "status":
		runStatus()
	default:
		printError("unknown command: %s", flag.Args()[0])
	}
}

func parseFlags() {
	apiKey = flag.StringP("key", "k", "", "sabnzbd api key")
	url = flag.StringP("url", "u", "https://localhost:9095", "sabnzbd url")
	verbose = flag.BoolP("verbose", "v", false, "verbose mode")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "sab <command>")
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.Parse()
}
