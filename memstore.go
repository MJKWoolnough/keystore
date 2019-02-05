package keystore

import (
	"io"
	"sort"
	"sync"

	"vimagination.zapto.org/memio"
)

type memStore struct {
	mu   sync.RWMutex
	data map[string]memio.Buffer
}

// NewMemStore creates a new memory-backed key-value store
func NewMemStore() Store {
	ms := new(memStore)
	ms.init()
	return ms
}

func (ms *memStore) init() {
	ms.data = make(map[string]memio.Buffer)
}

func (ms *memStore) Get(key string, r io.ReaderFrom) error {
	d := ms.get(key)
	if d == nil {
		return ErrUnknownKey
	}
	_, err := r.ReadFrom(&d)
	return err
}

func (ms *memStore) get(key string) memio.Buffer {
	ms.mu.RLock()
	d := ms.data[key]
	ms.mu.RUnlock()
	return d
}

func (ms *memStore) Set(key string, w io.WriterTo) error {
	d := make(memio.Buffer, 0)
	if _, err := w.WriteTo(&d); err != nil && err != io.EOF {
		return err
	}
	ms.set(key, d)
	return nil
}

func (ms *memStore) set(key string, d memio.Buffer) {
	ms.mu.Lock()
	ms.data[key] = d
	ms.mu.Unlock()
}

func (ms *memStore) Remove(key string) error {
	ms.mu.Lock()
	_, ok := ms.data[key]
	if !ok {
		ms.mu.Unlock()
		return ErrUnknownKey
	}
	delete(ms.data, key)
	ms.mu.Unlock()
	return nil
}

// Keys returns a sorted slice of all of the keys
func (ms *memStore) Keys() []string {
	ms.mu.RLock()
	s := make([]string, 0, len(ms.data))
	for key := range ms.data {
		s = append(s, key)
	}
	ms.mu.RUnlock()
	sort.Strings(s)
	return s
}
