package keystore

import (
	"io"

	"vimagination.zapto.org/memio"
)

// FileBackedMemStore combines both a FileStore and a MemStore
type FileBackedMemStore struct {
	FileStore
	memStore MemStore
}

// NewFileBackedMemStore create a new Store which uses the filesystem for
// permanent storage, but uses memory for caching
func NewFileBackedMemStore(baseDir, tmpDir string, mangler Mangler) (*FileBackedMemStore, error) {
	fs := new(FileBackedMemStore)
	if err := fs.init(baseDir, tmpDir, mangler); err != nil {
		return nil, err
	}
	fs.memStore.init()
	return fs, nil
}

// NewFileBackedMemStoreFromFileStore uses an existing FileStore to create a
// new File Backed Memory Store
func NewFileBackedMemStoreFromFileStore(filestore *FileStore) *FileBackedMemStore {
	fs := &FileBackedMemStore{
		FileStore: *filestore,
	}
	fs.memStore.init()
	return fs
}

// Get retrieves a key from the Store, first looking in the memcache and then
// going to the filesystem
func (fs *FileBackedMemStore) Get(key string, r io.ReaderFrom) error {
	err := fs.memStore.Get(key, r)
	if err == ErrUnknownKey {
		var buf memio.Buffer
		err = fs.FileStore.Get(key, &buf)
		if err == nil {
			fs.memStore.set(key, buf)
			_, err = r.ReadFrom(&buf)
		}
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
	if err = fs.FileStore.Set(key, &fbuf); err != nil {
		return err
	}
	fs.memStore.set(key, buf)
	return nil
}

// Remove deletes a key from both the memcache and the filesystem
func (fs *FileBackedMemStore) Remove(key string) error {
	if err := fs.FileStore.Remove(key); err != nil {
		return err
	}
	return fs.memStore.Remove(key)
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

// Rename moves data from an existing key to a new, unused key
func (fs *FileBackedMemStore) Rename(oldkey, newkey string) error {
	fs.memStore.mu.Lock()
	if err := fs.FileStore.Rename(oldkey, newkey); err != nil {
		fs.memStore.mu.Unlock()
		return err
	}
	delete(fs.memStore.data, oldkey)
	fs.memStore.mu.Unlock()
	return nil
}
