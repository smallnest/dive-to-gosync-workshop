package rwmap

import "sync"

type RWMap struct {
	sync.RWMutex
	m map[int]int
}

func NewRWMap(n int) *RWMap {
	return &RWMap{
		m: make(map[int]int, n),
	}
}
func (m *RWMap) Get(k int) (int, bool) {
	m.RLock()
	defer m.RUnlock()

	v, existed := m.m[k]
	return v, existed
}

func (m *RWMap) Set(k int, v int) {
	m.Lock()
	defer m.Unlock()

	m.m[k] = v
}

func (m *RWMap) Delete(k int) {
	m.Lock()
	defer m.Unlock()

	delete(m.m, k)
}

func (m *RWMap) Len() int {
	m.RLock()
	defer m.RUnlock()

	return len(m.m)
}

func (m *RWMap) Each(f func(k, v int) bool) {
	m.RLock()
	defer m.RUnlock()

	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}
}
