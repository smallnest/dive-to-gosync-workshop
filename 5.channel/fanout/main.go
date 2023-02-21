package main

import (
	"fmt"
	"reflect"
)

func fanOut(ch <-chan any, out []chan any, async bool) {
	go func() {
		defer func() {
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()

		for v := range ch {
			v := v
			for i := 0; i < len(out); i++ {
				i := i
				if async {
					go func() {
						out[i] <- v
					}()
				} else {
					out[i] <- v
				}
			}
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
		}

		for v := range ch {
			v := v
			for i := range cases {
				cases[i].Chan = reflect.ValueOf(out[i])
				cases[i].Send = reflect.ValueOf(v)
			}

			for _ = range cases { // for each channel
				chosen, _, _ := reflect.Select(cases)
				cases[chosen].Chan = reflect.ValueOf(nil)
			}
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
	source := asStream(nil)
	channels := make([]chan any, 5)

	fmt.Println("fanOut")
	for i := 0; i < 5; i++ {
		channels[i] = make(chan any)
	}
	fanOut(source, channels, false)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("channel#%d: %v\n", j, <-channels[j])
		}
	}

	fmt.Println("fanOut By Reflect")
	source = asStream(nil)
	for i := 0; i < 5; i++ {
		channels[i] = make(chan any)
	}
	fanOutReflect(source, channels)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("channel#%d: %v\n", j, <-channels[j])
		}
	}
}
