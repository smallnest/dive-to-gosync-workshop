package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

type SpinLock struct {
	f uint32
}

func (sl *SpinLock) Lock() {
	for !sl.TryLock() {
		runtime.Gosched()
	}
}

func (sl *SpinLock) Unlock() {
	atomic.StoreUint32(&sl.f, 0)
}

func (sl *SpinLock) TryLock() bool {
	return atomic.CompareAndSwapUint32(&sl.f, 0, 1)
}

func (sl *SpinLock) String() string {
	if atomic.LoadUint32(&sl.f) == 1 {
		return "Locked"
	}
	return "Unlocked"
}

func main() {
	var mu = &SpinLock{}
	go func() {
		mu.Lock()
		time.Sleep(time.Second)
		mu.Unlock()
	}()

	time.Sleep(time.Second)

	ok := mu.TryLock()
	if ok {
		fmt.Println("got the lock")
		// do something
		mu.Unlock()
		return
	} else {
		fmt.Println("can't get the lock")
	}

}
