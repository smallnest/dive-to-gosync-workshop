package main

import (
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/go-pkgz/syncs"
)

func main() {
	// 设置并发数为10,有限的goroutine数
	ewg := syncs.NewErrSizedGroup(10, syncs.Preemptive)
	var c uint32

	for i := 0; i < 1000; i++ {
		i := i
		ewg.Go(func() error {
			time.Sleep(time.Millisecond * 10)
			atomic.AddUint32(&c, 1)
			if i == 100 {
				return errors.New("err1")
			}
			if i == 200 {
				return errors.New("err2")
			}
			return nil
		})
	}

	// 等待任务完成
	if err := ewg.Wait(); err != nil {
		fmt.Printf("err: %v\n", err)
	}
	// 输出结果
	fmt.Println(c)
}
