package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/mgutz/ansi"
)

func colorizeMessage(color, prefix, message string, args ...interface{}) string {
	prefResult := ""
	if prefix != "" {
		prefResult = ansi.Color(prefix, color+"+b") + " " + ansi.ColorCode("reset")
	}
	return prefResult + ansi.Color(fmt.Sprintf(message, args...), color) + ansi.ColorCode("reset")
}

func listRec(a ...interface{}) {
	for i, x := range a {
		fmt.Printf("%v", x)
		if i+1 < len(a) {
			fmt.Print("\t")
		} else {
			fmt.Print("\n")
		}
	}
}

func needApiKey() {
	if *apiKey == "" {
		printError("need api key")
		os.Exit(2)
	}
}

func prettyFloat(f float64) string {
	return strconv.FormatFloat(float64(f), 'f', 1, 32)
}

func printError(message string, args ...interface{}) {
	log.Println(colorizeMessage("red", "error:", message, args...))
}

func printDebug(message string, args ...interface{}) {
	if *verbose {
		log.Println(fmt.Sprintf(message, args...))
	}
}

func printFatal(message string, args ...interface{}) {
	log.Fatal(colorizeMessage("red", "error:", message, args...))
}
