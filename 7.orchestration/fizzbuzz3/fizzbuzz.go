package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/marusama/cyclicbarrier"
)

// 编写一个可以从 1 到 n 输出代表这个数字的字符串的程序，要求：

// 	如果这个数字可以被 3 整除，输出 "fizz"。
// 	如果这个数字可以被 5 整除，输出 "buzz"。
// 	如果这个数字可以同时被 3 和 5 整除，输出 "fizzbuzz"。
// 	例如，当 n = 15，输出： 1, 2, fizz, 4, buzz, fizz, 7, 8, fizz, buzz, 11, fizz, 13, 14, fizzbuzz。

// 	假设有这么一个结构体：
// 	type FizzBuzz struct {}
// 	func (fb *FizzBuzz) fizz() {}
// 	func (fb *FizzBuzz) buzz() {}
// 	func (fb *FizzBuzz) fizzbuzz() {}
// 	func (fb *FizzBuzz) number() {}

// 请你实现一个有四个线程的多协程版 FizzBuzz，同一个 FizzBuzz 对象会被如下四个协程使用：

// 协程A将调用 fizz() 来判断是否能被 3 整除，如果可以，则输出 fizz。
// 协程B将调用 buzz() 来判断是否能被 5 整除，如果可以，则输出 buzz。
// 协程C将调用 fizzbuzz() 来判断是否同时能被 3 和 5 整除，如果可以，则输出 fizzbuzz。
// 协程D将调用 number() 来实现输出既不能被 3 整除也不能被 5 整除的数字。

type FizzBuzz struct {
	n int

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

	fb.wg.Wait()
}

func (fb *FizzBuzz) fizz() {
	defer fb.wg.Done()

	ctx := context.Background()
	v := 0
	for {
		fb.barrier.Await(ctx)
		v++

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
		}

	}

}
func (fb *FizzBuzz) buzz() {
	defer fb.wg.Done()

	ctx := context.Background()

	v := 0
	for {
		fb.barrier.Await(ctx)
		v++

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
		}
	}
}
func (fb *FizzBuzz) fizzbuzz() {
	defer fb.wg.Done()

	ctx := context.Background()
	v := 0
	for {
		fb.barrier.Await(ctx)
		v++

		if v > fb.n {
			return
		}

		if v%5 == 0 && v%3 == 0 {
			if v == fb.n {
				fmt.Print(" fizzbuzz。")
			} else {
				fmt.Print(" fizzbuzz,")
			}
		}
	}
}
func (fb *FizzBuzz) number() {
	defer fb.wg.Done()

	ctx := context.Background()
	v := 0
	for {
		fb.barrier.Await(ctx)
		v++

		if v > fb.n {
			return
		}

		if v%5 != 0 && v%3 != 0 {
			if v == fb.n {
				fmt.Printf(" %d。", v)
			} else {
				fmt.Printf(" %d,", v)
			}
		}

	}
}

func main() {
	fb := New(15)
	fb.start()
}
