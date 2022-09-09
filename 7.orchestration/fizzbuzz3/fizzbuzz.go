package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/marusama/cyclicbarrier"
)

type FizzBuzz struct {
	n int

	v       atomic.Int64
	barrier cyclicbarrier.CyclicBarrier
	wg      sync.WaitGroup
}

func New(n int) *FizzBuzz {
	return &FizzBuzz{
		n:       n,
		barrier: cyclicbarrier.New(4),
	}
}

func (fb *FizzBuzz) start() {
	fb.wg.Add(4)

	go fb.fizz()
	go fb.buzz()
	go fb.fizzbuzz()
	go fb.number()

	fb.v.Add(1)

	fb.wg.Wait()
}

func (fb *FizzBuzz) fizz() {
	defer fb.wg.Done()

	ctx := context.Background()
	for {
		fb.barrier.Await(ctx)

		v := int(fb.v.Load())

		if v > fb.n {
			return
		}

		if v%3 == 0 {
			if v%5 == 0 {
				continue
			}

			if v == fb.n {
				fmt.Print(" fizz。")
			} else {
				fmt.Print(" fizz,")
			}

			fb.v.Add(1)
		}

	}

}
func (fb *FizzBuzz) buzz() {
	defer fb.wg.Done()

	ctx := context.Background()
	for {
		fb.barrier.Await(ctx)

		v := int(fb.v.Load())

		if v > fb.n {
			return
		}

		if v%5 == 0 {
			if v%3 == 0 {
				continue
			}

			if v == fb.n {
				fmt.Print(" buzz。")
			} else {
				fmt.Print(" buzz,")
			}

			fb.v.Add(1)
		}
	}
}
func (fb *FizzBuzz) fizzbuzz() {
	defer fb.wg.Done()

	ctx := context.Background()
	for {
		fb.barrier.Await(ctx)

		v := int(fb.v.Load())

		if v > fb.n {
			return
		}

		if v%5 == 0 && v%3 == 0 {
			if v == fb.n {
				fmt.Print(" fizzbuzz。")
			} else {
				fmt.Print(" fizzbuzz,")
			}

			fb.v.Add(1)
			continue
		}
	}
}
func (fb *FizzBuzz) number() {
	defer fb.wg.Done()

	ctx := context.Background()
	for {
		fb.barrier.Await(ctx)

		v := int(fb.v.Load())

		if v > fb.n {
			return
		}

		if v%5 != 0 && v%3 != 0 {
			if v == fb.n {
				fmt.Printf(" %d。", v)
			} else {
				fmt.Printf(" %d,", v)
			}

			fb.v.Add(1)
			continue
		}

	}
}

func main() {
	fb := New(15)
	fb.start()
}
