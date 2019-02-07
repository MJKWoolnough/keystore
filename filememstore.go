package keystore

import (
	"io"

	"vimagination.zapto.org/memio"
)

// FileBackedMemStore combines both a FileStore and a MemStore
type FileBackedMemStore struct {
	fileStore FileStore
	memStore  MemStore
}

// NewFileBackedMemStore create a new Store which uses the filesystem for
// permanent storage, but uses memory for caching
func NewFileBackedMemStore(baseDir, tmpDir string, mangler Mangler) (*FileBackedMemStore, error) {
	fs := new(FileBackedMemStore)
	if err := fs.fileStore.init(baseDir, tmpDir, mangler); err != nil {
		return nil, err
	}
	fs.memStore.init()
	return fs, nil
}

// Get retrieves a key from the Store, first looking in the memcache and then
// going to the filesystem
func (fs *FileBackedMemStore) Get(key string, r io.ReaderFrom) error {
	err := fs.memStore.Get(key, r)
	if err == ErrUnknownKey {
		var buf memio.Buffer
		err = fs.fileStore.Get(key, &buf)
		if err != nil {
			return err
		}
		fs.memStore.set(key, buf)
		_, err = r.ReadFrom(&buf)
	}
	return err
}

// Set stores the key in both the memcache and the filesystem
func (fs *FileBackedMemStore) Set(key string, w io.WriterTo) error {
	var buf memio.Buffer
	_, err := w.WriteTo(&buf)
	if err != nil && err != io.EOF {
		return err
	}
	fbuf := buf
	if err = fs.fileStore.Set(key, &fbuf); err != nil {
		return err
	}
	fs.memStore.set(key, buf)
	return nil
}

// Remove deletes a key from both the memcache and the filesystem
func (fs *FileBackedMemStore) Remove(key string) error {
	if err := fs.fileStore.Remove(key); err != nil {
		return err
	}
	fs.memStore.Remove(key)
	return nil
}

// Clear removes keys from the memory cache. Specifying no keys removes all
// data
func (fs *FileBackedMemStore) Clear(keys ...string) {
	fs.memStore.mu.Lock()
	if len(keys) == 0 {
		for key := range fs.memStore.data {
			delete(fs.memStore.data, key)
		}
	} else {
		for _, key := range keys {
			delete(fs.memStore.data, key)
		}
	}
	fs.memStore.mu.Unlock()
}

// Keys returns a sorted slice of all of the keys
func (fs *FileBackedMemStore) Keys() []string {
	return fs.fileStore.Keys()
}
