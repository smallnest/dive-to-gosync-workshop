package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	var once sync.Once

	count := 0
	go func() {
		defer func() {
			count++
			recover()
		}()
		once.Do(func() {
			fmt.Println("exec Do")
			count = 1 / count
		})
	}()

	time.Sleep(time.Second)

	once.Do(func() {
		fmt.Println("exec here")
		count = 1 / count
	})

	fmt.Println("end")

	addr := "baidu.com"

	var conn net.Conn
	var err error

	once.Do(func() {
		conn, err = net.Dial("tcp", addr)
	})

	_ = conn
	_ = err
}
