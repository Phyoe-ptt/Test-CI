package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://time-service:8081/time")
		if err != nil {
			http.Error(w, "Failed to get time from time-service", http.StatusInternalServerError)
			log.Println("Error calling time-service:", err)
			return
		}
		defer resp.Body.Close()

		var timeString string
		_, err = fmt.Fscan(resp.Body, &timeString)
		if err != nil {
			http.Error(w, "Failed to parse time response", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Hello! The current time is: %s", timeString)
	})

	log.Println("Greeter service starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}