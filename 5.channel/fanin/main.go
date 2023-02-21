package main

import (
	"fmt"
	"reflect"
	"sync"
)

// https://github.com/campoy/justforfunc/blob/master/27-merging-chans/main.go

func fanIn(chans ...<-chan any) <-chan any {
	out := make(chan any)
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(chans))

		for _, c := range chans {
			go func(c <-chan any) {
				for v := range c {
					out <- v
				}
				wg.Done()
			}(c)
		}

		wg.Wait()
		close(out)
	}()
	return out
}

func fanInReflect(chans ...<-chan any) <-chan any {
	out := make(chan any)
	go func() {
		defer close(out)
		var cases []reflect.SelectCase
		for _, c := range chans {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		for len(cases) > 0 {
			i, v, ok := reflect.Select(cases)
			if !ok { //remove this case
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}
			out <- v.Interface()
		}
	}()
	return out

}

func fanInRec(chans ...<-chan any) <-chan any {
	switch len(chans) {
	case 0:
		c := make(chan any)
		close(c)
		return c
	case 1:
		return chans[0]
	case 2:
		return mergeTwo(chans[0], chans[1])
	default:
		m := len(chans) / 2
		return mergeTwo(
			fanInRec(chans[:m]...),
			fanInRec(chans[m:]...))
	}
}

func mergeTwo(a, b <-chan any) <-chan any {
	c := make(chan any)

	go func() {
		defer close(c)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					continue
				}
				c <- v
			}
		}
	}()
	return c
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
	fmt.Println("fanIn by goroutine:")
	done := make(chan struct{})
	ch := fanIn(asStream(done), asStream(done), asStream(done))
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("fanIn by reflect:")
	ch = fanInReflect(asStream(done), asStream(done), asStream(done))
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("fanIn by recursion:")
	ch = fanInRec(asStream(done), asStream(done), asStream(done))
	for v := range ch {
		fmt.Println(v)
	}
}
