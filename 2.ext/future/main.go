package main

import (
	"fmt"
	"time"

	"github.com/reugn/async"
)

func main() {
	p := async.NewPromise()
	go func() {
		time.Sleep(time.Millisecond * 100)
		p.Success(true)
	}()
	v, e := p.Future().Get()
	fmt.Println(v, e)

	p1 := async.NewPromise()
	p2 := async.NewPromise()
	p3 := async.NewPromise()
	go func() {
		time.Sleep(time.Millisecond * 100)
		p1.Success(1)
		time.Sleep(time.Millisecond * 200)
		p2.Success(2)
		time.Sleep(time.Millisecond * 300)
		p3.Success(3)
	}()

	v, e = async.FutureSeq([]async.Future{p1.Future(), p2.Future(), p3.Future()}).Get()
	fmt.Println(v, e)
}
