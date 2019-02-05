package keystore

import (
	"io"
	"sort"
	"sync"

	"vimagination.zapto.org/byteio"
	"vimagination.zapto.org/memio"
)

type memStore struct {
	mu   sync.RWMutex
	data map[string]memio.Buffer
}

type MemStore interface {
	Store
	io.WriterTo
	io.ReaderFrom
}

// NewMemStore creates a new memory-backed key-value store
func NewMemStore() MemStore {
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

func (ms *memStore) WriteTo(w io.Writer) (int64, error) {
	lw := byteio.StickyLittleEndianWriter{Writer: w}
	ms.mu.RLock()
	for key, value := range ms.data {
		lw.WriteStringX(key)
		lw.WriteUintX(uint64(len(value)))
		lw.Write(value)
	}
	ms.mu.RUnlock()
	return lw.Count, lw.Err
}

func (ms *memStore) ReadFrom(r io.Reader) (int64, error) {
	lr := byteio.StickyLittleEndianReader{Reader: r}
	ms.mu.Lock()
	for {
		key := lr.ReadStringX()
		buf := make(memio.Buffer, lr.ReadUintX())
		lr.Read(buf)
		if lr.Err != nil {
			break
		}
		ms.data[key] = buf
	}
	ms.mu.Unlock()
	return lr.Count, lr.Err
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
