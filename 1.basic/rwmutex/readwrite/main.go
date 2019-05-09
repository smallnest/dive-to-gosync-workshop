package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(12)

	var mu sync.RWMutex

	// 读锁
	go func() {
		defer wg.Done()

		mu.RLock()
		fmt.Println("before lock, reader0 RLock")
		time.Sleep(10 * time.Second)
		mu.RUnlock()
		fmt.Println("before lock,, reader0 RUnlock")
	}()

	time.Sleep(2 * time.Second)

	// 写锁
	go func() {
		defer wg.Done()

		mu.Lock()
		fmt.Println("writer1 Lock")
		time.Sleep(10 * time.Second)
		mu.Unlock()
		fmt.Println("writer1 Unlock")
	}()

	time.Sleep(time.Second)

	// 被前面的写锁 block
	for i := 1; i <= 10; i++ {
		i := i
		go func() {
			defer wg.Done()

			mu.RLock()
			fmt.Printf("after lock, reader %d RLock\n", i)
			time.Sleep(5 * time.Second)
			mu.RUnlock()
			fmt.Printf("after lock, reader %d RUnlock\n", i)
		}()
	}

	time.Sleep(time.Second)

	wg.Wait()
}
