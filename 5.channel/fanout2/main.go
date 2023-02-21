package main

import (
	"fmt"
	"reflect"
	"time"
)

func fanOut(ch <-chan any, out []chan any) {
	go func() {
		defer func() {
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()

		// roundrobin
		var i = 0
		var n = len(out)
		for v := range ch {
			v := v
			out[i] <- v
			i = (i + 1) % n
		}
	}()
}

func fanOutReflect(ch <-chan any, out []chan any) {
	go func() {
		defer func() {
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()

		cases := make([]reflect.SelectCase, len(out))
		for i := range cases {
			cases[i].Dir = reflect.SelectSend
			cases[i].Chan = reflect.ValueOf(out[i])

		}

		for v := range ch {
			v := v
			for i := range cases {
				cases[i].Send = reflect.ValueOf(v)
			}
			_, _, _ = reflect.Select(cases)
		}
	}()
}

func asStream(done <-chan struct{}) <-chan any {
	s := make(chan any)
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
	done := make(chan struct{})
	source := asStream(done)
	channels := make([]chan any, 5)

	fmt.Println("fanOut")
	for i := 0; i < 5; i++ {
		channels[i] = make(chan any)
	}
	fanOut(source, channels)
	for i := 0; i < 5; i++ {
		i := i
		go func() {
			for j := 0; j < 5; j++ {
				v, ok := <-channels[i]
				if ok {
					fmt.Printf("channel#%d: %v\n", i, v)
				}

			}
		}()
	}
	time.Sleep(time.Second)
	close(done)

	fmt.Println("fanOut By Reflect")
	done = make(chan struct{})
	source = asStream(done)
	for i := 0; i < 5; i++ {
		channels[i] = make(chan any)
	}
	fanOutReflect(source, channels)
	for i := 0; i < 5; i++ {
		i := i
		go func() {
			for j := 0; j < 5; j++ {
				v, ok := <-channels[i]
				if ok {
					fmt.Printf("channel#%d: %v\n", i, v)
				}
			}
		}()
	}
	time.Sleep(time.Second)
	close(done)

}
