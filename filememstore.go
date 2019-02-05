package keystore

import (
	"io"

	"vimagination.zapto.org/memio"
)

type fileBackedMemStore struct {
	fileStore
	memStore
}

// NewFileBackedMemStore create s new Store which uses the filesystem for
// permanent storage, but uses memory for caching
func NewFileBackedMemStore(baseDir, tmpDir string) (Store, error) {
	fs := new(fileBackedMemStore)
	if err := fs.fileStore.init(baseDir, tmpDir); err != nil {
		return nil, err
	}
	fs.memStore.init()
	return fs, nil
}

func (fs *fileBackedMemStore) Get(key string, r io.ReaderFrom) error {
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

func (fs *fileBackedMemStore) Set(key string, w io.WriterTo) error {
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

func (fs *fileBackedMemStore) Remove(key string) error {
	if err := fs.fileStore.Remove(key); err != nil {
		return err
	}
	fs.memStore.Remove(key)
	return nil
}

// Keys returns a sorted slice of all of the keys
func (fs *fileBackedMemStore) Keys() []string {
	return fs.fileStore.Keys()
}
