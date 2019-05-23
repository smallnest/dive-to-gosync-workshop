package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var once sync.Once

	var count = 0
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
}
