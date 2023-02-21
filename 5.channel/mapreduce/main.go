package main

import "fmt"

func mapChan[T, K any](in <-chan T, fn func(T) K) <-chan K {
	out := make(chan K)
	if in == nil {
		close(out)
		return out
	}

	go func() {
		defer close(out)

		for v := range in {
			out <- fn(v)
		}
	}()

	return out
}

func reduce[T, K any](in <-chan T, fn func(r K, v T) K) K {
	var out K

	if in == nil {
		return out
	}

	for v := range in {
		out = fn(out, v)
	}

	return out
}

func asStream(done <-chan struct{}) <-chan int {
	s := make(chan int)
	values := []int{1, 2, 3, 4, 5}
	go func() {
		defer close(s)

		for _, v := range values {
			select {
			case <-done:
				return
			case s <- v:
			}
		}

	}()
	return s
}

func main() {
	in := asStream(nil)

	// map op: time 10
	mapFn := func(v int) int {
		return v * 10
	}

	// reduce op: sum
	reduceFn := func(r, v int) int {
		return r + v
	}

	sum := reduce(mapChan(in, mapFn), reduceFn)
	fmt.Println(sum)
}
