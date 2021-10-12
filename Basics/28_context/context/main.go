package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	minLatency = 10
	maxLatency = 5000
)

func main() {
	// cancellation signal between coroutines
	// fmt.Println("Starting to search for nyc-location....")
	res, err := Search("nyc", "london")
	if err!=nil {
		panic(err)
	}
	fmt.Println("got results:", res)
	
}

func Search(from , to string) ([]string, error) {

	rand.Seed(time.Now().Unix())
	
	latency := rand.Intn(maxLatency - minLatency + 1) - minLatency
	fmt.Printf("started to search for %s - %s takes %d ms...\n", from, to , latency)
	time.Sleep(time.Duration(latency) * time.Millisecond)
	return []string{from + "-" + to + "-british airways-11am", from + "-" + to + "-delta airways-12am"}, nil
}