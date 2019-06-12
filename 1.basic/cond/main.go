package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// https://stackoverflow.com/questions/51371587/broadcast-in-golang

func main() {
	var m sync.Mutex
	c := sync.NewCond(&m)

	ready := make(chan struct{}, 10)
	isReady := false

	for i := 0; i < 10; i++ {
		i := i
		go func() {
			m.Lock()

			time.Sleep(time.Duration(rand.Int63n(2)) * time.Second)

			ready <- struct{}{} // 运动员i准备就绪
			for !isReady {
				c.Wait()
			}
			log.Printf("%d started\n", i)
			m.Unlock()
		}()
	}

	// false broadcast
	c.Broadcast()

	// 裁判员检查所有的运动员是否就绪
	for i := 0; i < 10; i++ {
		<-ready
	}

	// 运动员都已准备就绪，发令枪响, broadcast
	// m.Lock()
	isReady = true
	c.Broadcast()
	// m.Unlock()

	time.Sleep(time.Second)
}
