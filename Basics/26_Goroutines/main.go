/*
			-------Go Routines--------
1. Lightweight thread
2. Mapped onto OS thread(s) by the Go runtime
3. 2KB intial stack size, grows dynamically
4. Safe to spawn hundreds of thousands of goroutines
5. Can communicate via Go-channels


      ---------WaitGroup------------
1. A synchronization mechanism
2. Goroutine can be added to WaitGroup
3. GR informs WG when it gas finished execution
4. WG blocks calling GR until all registered GRs are done
*/

package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func sendRequest(url string) {
	defer wg.Done()
	res, err := http.Get(url) // send Request
	if err != nil {
		panic(err) // end progr  of error
	}
	
	fmt.Printf("[%d] %s\n", res.StatusCode, url)
}

func main() {
	var urls []string = []string {
		"google.com", "facebook.com", "twitter.com", "youtube.com", "github.com", "linkedin.com", 
		"gitlab.com", "yahoo.com",
	}
		
	for _, url := range urls {
		go sendRequest("https://"+url)
		wg.Add(1)
	}
	
	// wg.Wait()
}
