package syncmap

import "sync"

type IntMap struct {
	m map[int]int
	sync.Mutex
}

func (m *IntMap) Get(k int) (v int, ok bool) {
	m.Lock()
	defer m.Unlock()

	v, ok = m.m[k]
	return v, ok
}

func (m *IntMap) Set(k, v int) {
	m.Lock()
	defer m.Unlock()

	m.m[k] = v
}

func (m *IntMap) Len() int {
	m.Lock()
	defer m.Unlock()

	return len(m.m)
}

func (m *IntMap) Range(fn func(k, v int) error) error {
	m.Lock()
	defer m.Unlock()

	var err error
	for k, v := range m.m {
		err = fn(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}
