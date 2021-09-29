package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock sync.Mutex
	rwLock sync.RWMutex
	count int
)

func basics() {
	iterations := 1000
	for i := 0; i < iterations; i++ {
		go increment()
	}
	
	time.Sleep(1 * time.Second)
	
	fmt.Println("Resulted Count is:", count)
}

	
func readAndWrite() {
	go read()
	go write()
	
	time.Sleep(5* time.Second)
	fmt.Println("Done")
}

func read() {
	rwLock.Lock()
	defer rwLock.Unlock()
	
	fmt.Println("Read locking...")
	time.Sleep(1*time.Second)
	fmt.Println("Read unlocking...")
}

func write() {
	rwLock.Lock()
	defer rwLock.Unlock()
	
	fmt.Println("Write locking...")
	time.Sleep(1*time.Second)
	fmt.Println("Write unlocking...")
}

func main() {
	basics()
	readAndWrite()
}

func increment() {
	lock.Lock()
	count++
	lock.Unlock()
}