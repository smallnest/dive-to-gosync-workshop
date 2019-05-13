package main

import (
	"sync"
	"time"
)

var mu sync.Mutex

func main() {
	go lock()
	time.Sleep(1e9)
}

func lock() {
	mu.Lock()
	defer mu.Lock()
}
