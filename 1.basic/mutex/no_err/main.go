package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Mutex
	fmt.Print("A, ")
	m.Lock()

	go func() {
		time.Sleep(200 * time.Millisecond)
		m.Unlock()
	}()

	// 等另一个goroutine释放锁之后，main goroutine就可以获取到锁
	m.Lock()
	fmt.Print("B ")
}
