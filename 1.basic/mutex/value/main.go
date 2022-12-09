package main

import "sync"

// https://github.com/carlmjohnson/syncx/blob/main/mutex.go

type Mutex[T any] struct {
	mu    sync.RWMutex
	value T
}

func NewMutex[T any](initial T) *Mutex[T] {
	var m Mutex[T]
	m.value = initial
	return &m
}

func (m *Mutex[T]) ReadLock(f func(value T)) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	f(m.value)
}

func (m *Mutex[T]) Lock(f func(value *T)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	value := m.value
	f(&value)
	m.value = value
}

func (m *Mutex[T]) Load() T {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.value
}

func (m *Mutex[T]) Store(value T) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.value = value
}
