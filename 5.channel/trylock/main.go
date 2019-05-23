package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
)

// defined in sync.Mutex
const mutexLocked = 1 << iota

type Mutex struct {
	mu sync.Mutex
}

func (m *Mutex) Lock() {
	m.mu.Lock()
}

func (m *Mutex) Unlock() {
	m.mu.Unlock()
}

func (m *Mutex) TryLock() bool {
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.mu)), 0, mutexLocked)
}

func (m *Mutex) IsLocked() bool {
	return atomic.LoadInt32((*int32)(unsafe.Pointer(&m.mu))) == mutexLocked
}

func main() {
	var m Mutex
	if m.TryLock() {
		fmt.Println("locked: ", m.IsLocked())
		m.Unlock()
	} else {
		fmt.Println("failed to lock")
	}
}
