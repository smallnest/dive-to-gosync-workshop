package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/go-pkgz/syncs"
)

func main() {
	// 设置并发数是10
	swg := syncs.NewSizedGroup(10, syncs.Preemptive)
	// swg := syncs.NewSizedGroup(10, syncs.Preemptive)
	var c uint32

	// 执行1000个子任务，只会有10个goroutine去执行
	for i := 0; i < 1000; i++ {
		swg.Go(func(ctx context.Context) {
			time.Sleep(5 * time.Millisecond)
			atomic.AddUint32(&c, 1)
		})
	}

	// 等待任务完成
	swg.Wait()
	// 输出结果
	fmt.Println(c)
}
