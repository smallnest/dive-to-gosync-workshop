package main

import (
	"fmt"
	"unsafe"
)

type waitGroup1 struct {
	state1 [3]uint32
}

type waitGroup2 struct {
	state1 uint64
	sema   uint32
}

type waitGroup3 struct {
	state1 uint64
	sema   uint32
}

type A1 struct {
	a int32
	waitGroup1
}

type A2 struct {
	a int32
	waitGroup2
}

type A3 struct {
	a int32
	waitGroup3
}

func main() {
	var wg1 waitGroup1
	var a1 A1
	fmt.Println(unsafe.Alignof(wg1.state1))
	fmt.Println(unsafe.Pointer(&wg1.state1))
	fmt.Println(uintptr(unsafe.Pointer(&wg1.state1))%8 == 0)
	fmt.Println(uintptr(unsafe.Pointer(&a1.waitGroup1.state1))%8 == 0)
	fmt.Println()

	var wg2 waitGroup2
	var a2 A2
	fmt.Println(unsafe.Alignof(wg2.state1))
	fmt.Println(unsafe.Pointer(&wg2.state1))
	fmt.Println(uintptr(unsafe.Pointer(&wg2.state1))%8 == 0)
	fmt.Println(uintptr(unsafe.Pointer(&a2.waitGroup2.state1))%8 == 0)
	fmt.Println()

	var wg3 waitGroup3
	var a3 A3
	fmt.Println(unsafe.Alignof(wg3.state1))
	fmt.Println(unsafe.Pointer(&wg3.state1))
	fmt.Println(uintptr(unsafe.Pointer(&wg3.state1))%8 == 0)
	fmt.Println(uintptr(unsafe.Pointer(&a3.waitGroup3.state1))%8 == 0)
}
