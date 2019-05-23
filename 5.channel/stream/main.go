package main

import (
	"fmt"
)

func asStream(done <-chan struct{}, values ...interface{}) <-chan interface{} {
	s := make(chan interface{})
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
func asRepeatStream(done <-chan struct{}, values ...interface{}) <-chan interface{} {
	s := make(chan interface{})
	go func() {
		defer close(s)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case s <- v:
				}
			}
		}
	}()
	return s
}

func takeN(done <-chan struct{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func takeFn(done <-chan struct{}, valueStream <-chan interface{}, fn func(interface{}) bool) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for {
			select {
			case <-done:
				return
			case v := <-valueStream:
				if fn(v) {
					takeStream <- v
				}
			}
		}
	}()
	return takeStream
}

func takeWhile(done <-chan struct{}, valueStream <-chan interface{}, fn func(interface{}) bool) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for {
			select {
			case <-done:
				return
			case v := <-valueStream:
				if !fn(v) {
					return
				}
				takeStream <- v
			}
		}
	}()
	return takeStream
}

func skipN(done <-chan struct{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case <-valueStream:
			}
		}
		for {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()

	return takeStream
}

func skipFn(done <-chan struct{}, valueStream <-chan interface{}, fn func(interface{}) bool) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for {
			select {
			case <-done:
				return
			case v := <-valueStream:
				if !fn(v) {
					takeStream <- v
				}
			}
		}
	}()
	return takeStream
}

func skipWhile(done <-chan struct{}, valueStream <-chan interface{}, fn func(interface{}) bool) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		take := false
		for {
			select {
			case <-done:
				return
			case v := <-valueStream:
				if !take {
					take = !fn(v)
					if !take {
						continue
					}
				}
				takeStream <- v
			}
		}
	}()
	return takeStream
}

func main() {
	fmt.Println("asStream:")
	done := make(chan struct{})
	ch := asStream(done, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("asRepeatStream:")
	done = make(chan struct{})
	ch = asRepeatStream(done, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	i := 0
	for v := range ch {
		fmt.Println(v)
		i++
		if i == 20 {
			break
		}
	}

	fmt.Println("takeN:")
	done = make(chan struct{})
	for v := range takeN(done, asRepeatStream(done, 1, 2, 3), 6) {
		fmt.Println(v)
	}

	evenFn := func(v interface{}) bool {
		return v.(int)%2 == 0
	}
	lessFn := func(v interface{}) bool {
		return v.(int) < 3
	}

	fmt.Println("takeFn:")
	done = make(chan struct{})
	i = 0
	for v := range takeFn(done, asRepeatStream(done, 1, 2, 3), evenFn) {
		fmt.Println(v)
		i++
		if i == 20 {
			break
		}
	}

	fmt.Println("takeWhile:")
	done = make(chan struct{})
	for v := range takeWhile(done, asRepeatStream(done, 1, 2, 3), lessFn) {
		fmt.Println(v)
	}

	fmt.Println("skipN:")
	done = make(chan struct{})
	for v := range takeN(done, skipN(done, asRepeatStream(done, 1, 2, 3), 2), 4) {
		fmt.Println(v)
	}

	fmt.Println("skipFn:")
	done = make(chan struct{})
	for v := range takeN(done, skipFn(done, asRepeatStream(done, 1, 2, 3), evenFn), 4) {
		fmt.Println(v)
	}

	fmt.Println("skipWhile:")
	done = make(chan struct{})
	for v := range takeN(done, skipWhile(done, asRepeatStream(done, 1, 2, 3), lessFn), 4) {
		fmt.Println(v)
	}
}
