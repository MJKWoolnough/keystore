package keystore

import (
	"io"

	"vimagination.zapto.org/memio"
)

type FileBackedMemStore struct {
	fileStore FileStore
	memStore  MemStore
}

// NewFileBackedMemStore create s new Store which uses the filesystem for
// permanent storage, but uses memory for caching
func NewFileBackedMemStore(baseDir, tmpDir string, mangler Mangler) (*FileBackedMemStore, error) {
	fs := new(FileBackedMemStore)
	if err := fs.fileStore.init(baseDir, tmpDir, mangler); err != nil {
		return nil, err
	}
	fs.memStore.init()
	return fs, nil
}

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

func (fs *FileBackedMemStore) Remove(key string) error {
	if err := fs.fileStore.Remove(key); err != nil {
		return err
	}
	fs.memStore.Remove(key)
	return nil
}

// Keys returns a sorted slice of all of the keys
func (fs *FileBackedMemStore) Keys() []string {
	return fs.fileStore.Keys()
}
