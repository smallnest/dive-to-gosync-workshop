package chans

import (
	"context"
	"testing"
)

func BenchmarkChan_WithTwoCases(b *testing.B) {
	ch := make(chan struct{}, 1000)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ch:
			}
		}
	}()

	for i := 0; i < b.N; i++ {
		ch <- struct{}{}
	}
	close(done)
}

func BenchmarkChan_WithOneCase(b *testing.B) {
	ch := make(chan struct{}, 1000)
	go func() {
		for {
			select {
			case <-ch:
			}
		}
	}()

	for i := 0; i < b.N; i++ {
		ch <- struct{}{}
	}
}

func BenchmarkChan_WithContext(b *testing.B) {
	ch := make(chan struct{}, 1000)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		done := ctx.Done()
		for {
			select {
			case <-done:
			case <-ch:
			}
		}
	}()

	for i := 0; i < b.N; i++ {
		ch <- struct{}{}
	}

	cancel()
}

func BenchmarkChan_Range(b *testing.B) {
	ch := make(chan struct{}, 1000)
	go func() {
		for range ch {
		}
	}()

	for i := 0; i < b.N; i++ {
		ch <- struct{}{}
	}
}

func BenchmarkChan_WithoutCase(b *testing.B) {
	ch := make(chan struct{}, 1000)
	go func() {
		for {
			<-ch
		}
	}()

	for i := 0; i < b.N; i++ {
		ch <- struct{}{}
	}
}
