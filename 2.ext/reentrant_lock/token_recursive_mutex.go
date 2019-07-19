package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type TokenRecursiveMutex struct {
	sync.Mutex
	token     int64
	recursion int32
}

func (m *TokenRecursiveMutex) Lock(token int64) {
	if atomic.LoadInt64(&m.token) == token {
		m.recursion++
		return
	}

	m.Mutex.Lock()
	// we are now inside the lock
	atomic.StoreInt64(&m.token, token)
	m.recursion = 1
}

func (m *TokenRecursiveMutex) Unlock(token int64) {
	if atomic.LoadInt64(&m.token) != token {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.token, token))
	}

	m.recursion--
	if m.recursion != 0 {
		return
	}

	atomic.StoreInt64(&m.token, 0)
	m.Mutex.Unlock()
}
