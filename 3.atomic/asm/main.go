package test

import (
	"sync/atomic"
)

func add(i *int64) {
	atomic.AddInt64(i, 100) //
}

func cas(i *int64) {
	atomic.CompareAndSwapInt64(i, 0, 100)
}

func load(i *int64) {
	atomic.LoadInt64(i)
}

func store(i *int64) {
	atomic.StoreInt64(i, 100)
}

func swap(i *int64) {
	atomic.SwapInt64(i, 100)
}
