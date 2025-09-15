package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now().Format(time.RFC850)
		fmt.Fprintln(w, currentTime)
	})

	log.Println("Time service starting on :8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}