package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GC()
	mem := &runtime.MemStats{}
	runtime.ReadMemStats(mem)
	fmt.Println("before:", mem.HeapInuse)
	for i := 0; i < 1000; i++ {
		task()
	}
	runtime.GC()
	runtime.ReadMemStats(mem)
	fmt.Println("after:", mem.HeapInuse)

	time.Sleep(15 * time.Second)
	runtime.GC()
	runtime.ReadMemStats(mem)
	fmt.Println("fired:", mem.HeapInuse)

}

func task() {
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("timeout")
		return
	default:
	}
}
