package main

import "sync"

func main() {
	var m = make(map[int]int, 10)

	go func() {
		for {
			m[1] = 1
		}
	}()

	go func() {
		for {
			_ = m[2]
		}
	}()

	select {}

	var mm sync.Map
}
