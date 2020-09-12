package main

import (
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})

	var ready uint32

	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

			// 运动员i准备就绪
			atomic.AddUint32(&ready, 1)
			log.Printf("运动员#%d 已准备就绪\n", i)

			c.Broadcast()
		}(i)
	}

	for atomic.LoadUint32(&ready) != 10 {
		c.L.Lock()
		c.Wait()
		c.L.Unlock()
	}

	//所有的运动员是否就绪
	log.Println("所有运动员都准备就绪。比赛开始，3，2，1, ......")
}
