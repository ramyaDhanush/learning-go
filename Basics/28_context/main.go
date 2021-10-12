package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const (
	minLatency = 10
	maxLatency = 5000
	timeOut = 3000
)

func main() {
	// cancellation signal between coroutines
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut) * time.Millisecond)
	defer cancel()
	
	fmt.Println(ctx)
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