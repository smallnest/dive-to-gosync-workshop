package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

// https://texlution.com/post/golang-lock-free-values-with-atomic-value/

type Config struct {
	sync.RWMutex
	endpoint string
}

func BenchmarkMutexSet(b *testing.B) {
	config := Config{}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			config.Lock()
			config.endpoint = "golang.org"
			config.Unlock()
		}
	})
}

func BenchmarkMutexGet(b *testing.B) {
	config := Config{endpoint: "golang.org"}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			config.RLock()
			_ = config.endpoint
			config.RUnlock()
		}
	})
}

func BenchmarkAtomicSet(b *testing.B) {
	var config atomic.Value
	c := Config{endpoint: "golang.org"}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			config.Store(c)
		}
	})
}

func BenchmarkAtomicGet(b *testing.B) {
	var config atomic.Value
	config.Store(Config{endpoint: "golang.org"})
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = config.Load().(Config)
		}
	})
}
