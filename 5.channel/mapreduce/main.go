package main

import "fmt"

func mapChan(in <-chan interface{}, fn func(interface{}) interface{}) <-chan interface{} {
	out := make(chan interface{})
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

func reduce(in <-chan interface{}, fn func(r, v interface{}) interface{}) interface{} {
	if in == nil {
		return nil
	}

	out := <-in
	for v := range in {
		out = fn(out, v)
	}

	return out
}

func asStream(done <-chan struct{}) <-chan interface{} {
	s := make(chan interface{})
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
	mapFn := func(v interface{}) interface{} {
		return v.(int) * 10
	}

	// reduce op: sum
	reduceFn := func(r, v interface{}) interface{} {
		return r.(int) + v.(int)
	}

	sum := reduce(mapChan(in, mapFn), reduceFn)
	fmt.Println(sum)
}
