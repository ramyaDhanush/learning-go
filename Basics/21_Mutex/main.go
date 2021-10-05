package main

import (
	"fmt"
	"sync"
	"time"
)

/*
----MUTEX------
1. Synchronization mechanism
2. Mutually Exclusive Lock
3. synchronizes access to shared resources
4. Goroutines can request a Mutex lock on a shared resource
5. Other Goroutines cannot access the resource until mutex is unlocked
*/
var (
	lock sync.Mutex
	count int
)

func main() {
	iterations := 1000
	for i := 0; i < iterations; i++ {
		go increment()
	}
	
	time.Sleep(1 * time.Second)
	
	fmt.Println("Resulted Count is:", count)
}

func increment() {
	lock.Lock()
	count++
	lock.Unlock()
}