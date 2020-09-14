package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	once.Do(func() {
		once.Do(func() {
			fmt.Println("初始化")
		})
	})
}
