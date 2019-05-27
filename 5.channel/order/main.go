package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	var i = 0

	for ; i < 10; i++ {
		ch <- i
	}

	go func() {
		for {
			i++
			ch <- i
		}
	}()

	for j := 0; j < 50; j++ {
		time.Sleep(time.Second)
		k := <-ch // 是优先读取缓存的数据，还是直接读取阻塞住的gorontine发送的数据
		fmt.Println(k)
	}
}
