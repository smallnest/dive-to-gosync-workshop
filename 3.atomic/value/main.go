package main

import (
	"fmt"
	"sync/atomic"
)

type T struct{ a int }

func main() {
	var tv atomic.Value

	var ta, tb T
	// store
	tv.Store(ta)

	// load
	tv1 := tv.Load().(T)
	fmt.Println(tv1 == ta) // true

	// store another
	tv.Store(tb)
	tv2 := tv.Load().(T)
	fmt.Println(tv2 == tb) // true
}
