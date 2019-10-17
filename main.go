package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var startDate duration
var dateSince duration

var mutex = &async.Mutex{}

func main() {
	startDate = time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	dateSince = time.Since(start)
	fmt.Println(since)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Time since last pizza:", since.String())
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
