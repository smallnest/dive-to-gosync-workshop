package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type T struct{ a int }

func main() {
	ppt()
}

func ppt() {
	var pT *T

	var unsafePPT = (*unsafe.Pointer)(unsafe.Pointer(&pT))

	var ta, tb T
	// store
	atomic.StorePointer(unsafePPT, unsafe.Pointer(&ta))

	// load
	pa1 := (*T)(atomic.LoadPointer(unsafePPT))
	fmt.Println(pa1 == &ta) // true

	// swap
	pa2 := atomic.SwapPointer(unsafePPT, unsafe.Pointer(&tb))
	fmt.Println((*T)(pa2) == &ta) // true

	// compare and swap
	b := atomic.CompareAndSwapPointer(unsafePPT, pa2, unsafe.Pointer(&tb))
	fmt.Println(b) // false
	b = atomic.CompareAndSwapPointer(unsafePPT, unsafe.Pointer(&tb), pa2)
	fmt.Println(b) // true
}
