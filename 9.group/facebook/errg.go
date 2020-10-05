package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/facebookgo/errgroup"
)

func main() {
	var g errgroup.Group
	g.Add(3)

	// 启动第一个子任务,它执行成功
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("exec #1")
		g.Done()
	}()

	// 启动第二个子任务，它执行失败
	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("exec #2")
		g.Error(errors.New("failed to exec #2"))
		g.Done()
	}()

	// 启动第三个子任务，它执行成功
	go func() {
		time.Sleep(15 * time.Second)
		fmt.Println("exec #3")
		g.Done()
	}()

	// 等待所有的goroutine完成，并检查error
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully exec all")
	} else {
		fmt.Println("failed:", err)
	}
}
