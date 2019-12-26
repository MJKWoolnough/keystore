package keystore

import (
	"io"
	"sort"
	"sync"

	"vimagination.zapto.org/byteio"
	"vimagination.zapto.org/memio"
)

// MemStore implements Store and does so entirely in memory
type MemStore struct {
	mu   sync.RWMutex
	data map[string]memio.Buffer
}

// NewMemStore creates a new memory-backed key-value store
func NewMemStore() *MemStore {
	ms := new(MemStore)
	ms.init()
	return ms
}

func (ms *MemStore) init() {
	ms.data = make(map[string]memio.Buffer)
}

// Get retrieves the key data from memory
func (ms *MemStore) Get(key string, r io.ReaderFrom) error {
	d := ms.get(key)
	if d == nil {
		return ErrUnknownKey
	}
	_, err := r.ReadFrom(&d)
	return err
}

// GetAll retrieves data for all of the keys given. Useful to reduce locking.
// Unknown Key errors are not returned, only errors from the ReaderFrom's
func (ms *MemStore) GetAll(data map[string]io.ReaderFrom) error {
	var err error
	ms.mu.RLock()
	for k, d := range data {
		buf, ok := ms.data[k]
		if !ok {
			continue
		}
		if _, err = d.ReadFrom(&buf); err != nil {
			return err
		}
	}
	ms.mu.RUnlock()
	return err
}

func (ms *MemStore) get(key string) memio.Buffer {
	ms.mu.RLock()
	d := ms.data[key]
	ms.mu.RUnlock()
	return d
}

// Set stores the key data in memory
func (ms *MemStore) Set(key string, w io.WriterTo) error {
	d := make(memio.Buffer, 0)
	if _, err := w.WriteTo(&d); err != nil && err != io.EOF {
		return err
	}
	ms.set(key, d)
	return nil
}

// SetAll set data for all of the keys given. Useful to reduce locking.
// Will return the first error found, so may not set all data.
func (ms *MemStore) SetAll(data map[string]io.WriterTo) error {
	var err error
	ms.mu.Lock()
	for k, d := range data {
		var buf memio.Buffer
		if _, err = d.WriteTo(&buf); err != nil {
			break
		}
		ms.data[k] = buf
	}
	ms.mu.Unlock()
	return err
}

func (ms *MemStore) set(key string, d memio.Buffer) {
	ms.mu.Lock()
	ms.data[key] = d
	ms.mu.Unlock()
}

// Remove deletes the key data from memory
func (ms *MemStore) Remove(key string) error {
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

// RemoveAll will attempt to remove all keys given. It does not return an error
// if a key doesn't exist
func (ms *MemStore) RemoveAll(keys ...string) {
	ms.mu.Lock()
	for _, key := range keys {
		delete(ms.data, key)
	}
	ms.mu.Unlock()
}

// WriteTo implements the io.WriterTo interface allowing a MemStore to be
// be stored in another Store
func (ms *MemStore) WriteTo(w io.Writer) (int64, error) {
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

// ReadFrom implements the io.ReaderFrom interface allowing a MemStore to be
// be retrieved in another Store
func (ms *MemStore) ReadFrom(r io.Reader) (int64, error) {
	lr := byteio.StickyLittleEndianReader{Reader: r}
	ms.mu.Lock()
	for {
		key := lr.ReadStringX()
		if lr.Err == io.EOF {
			lr.Err = nil
			break
		}
		buf := make(memio.Buffer, lr.ReadUintX())
		lr.Read(buf)
		if lr.Err != nil {
			if lr.Err == io.EOF {
				lr.Err = io.ErrUnexpectedEOF
			}
			break
		}
		ms.data[key] = buf
	}
	ms.mu.Unlock()
	return lr.Count, lr.Err
}

// Keys returns a sorted slice of all of the keys
func (ms *MemStore) Keys() []string {
	ms.mu.RLock()
	s := make([]string, 0, len(ms.data))
	for key := range ms.data {
		s = append(s, key)
	}
	ms.mu.RUnlock()
	sort.Strings(s)
	return s
}

// Exists returns true when the key exists within the store
func (ms *MemStore) Exists(key string) bool {
	ms.mu.RLock()
	_, ok := ms.data[key]
	ms.mu.RUnlock()
	return ok
}

// Rename moves data from an existing key to a new, unused key
func (ms *MemStore) Rename(oldkey, newkey string) error {
	ms.mu.Lock()
	var err error
	if d, ok := ms.data[oldkey]; !ok {
		err = ErrUnknownKey
	} else if _, ok = ms.data[newkey]; ok {
		err = ErrKeyExists
	} else {
		ms.data[newkey] = d
		delete(ms.data, oldkey)
	}
	ms.mu.Unlock()
	return err
}
