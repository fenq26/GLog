package main

import (
	"fmt"
	"strings"
)

var window []Log
var maxWindow = 100

func analyze(logData Log) {
	window = append(window, logData)

	if len(window) > maxWindow {
		window = window[1:]
	}

	errorCount := 0
	timeoutCount := 0

	for _, := range window {
		msg := strings.ToLower(l.Message)

		if strings.Contains(msg, "timeout") {
			timeoutCount++
		}
	}

	score := errorCount*2 + timeoutCount*3

	if score > 15 {
		fmt.Println("Anomaly code detected")
		fmt.Println("User:", logData.Owner)
		fmt.Println("Service:", logData.Service)
		fmt.Println("Score:", score)
	}
}