package main

import (
	"fmt"
	"net/http"
)

func worker() {
	for logData := range LogBuffer {
		analyze(logData)
	}
}

func main() {
	for i := 0; i < 3; i++ {
		go worker()
	} 

	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/logs", logHandler)

	fmt.Println("GLog running on local (:8080)" )
	http.listenAndServe(":8080", nil)
}