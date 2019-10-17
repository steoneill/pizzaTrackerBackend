package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var startDate time.Time
var dateSince time.Duration

func main() {
	startDate = time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	dateSince = time.Since(startDate)
	fmt.Println(dateSince)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Time since last pizza:", dateSince)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
