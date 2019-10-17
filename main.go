package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	since := time.Since(start)
	fmt.Println(since)
}



