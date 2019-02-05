package keystore

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"vimagination.zapto.org/errors"
)

type fileStore struct {
	baseDir, tmpDir string
}

// NewFileStore creates a file backed key-value store
func NewFileStore(baseDir, tmpDir string) (Store, error) {
	fs := new(fileStore)
	if err := fs.init(baseDir, tmpDir); err != nil {
		return nil, err
	}
	return fs, nil
}

func (fs *fileStore) init(baseDir, tmpDir string) error {
	fs.baseDir = baseDir
	fs.tmpDir = tmpDir
	return nil
}

func (fs *fileStore) Get(key string, r io.ReaderFrom) error {
	if err := testKey(key); err != nil {
		return err
	}
	f, err := os.Open(filepath.Join(fs.baseDir, key))
	if err != nil {
		if os.IsNotExist(err) {
			return ErrUnknownKey
		}
		return errors.WithContext("error opening key file: ", err)
	}
	_, err = r.ReadFrom(f)
	f.Close()
	return err
}

func (fs *fileStore) Set(key string, w io.WriterTo) error {
	err := testKey(key)
	if err != nil {
		return err
	}
	var f *os.File
	if fs.tmpDir != "" {
		f, err = ioutil.TempFile(fs.tmpDir, "keystore")
	} else {
		f, err = os.Create(filepath.Join(fs.baseDir, key))
	}
	if err != nil {
		return errors.WithContext("error opening file for writing: ", err)
	}
	if _, err = w.WriteTo(f); err != nil {
		f.Close()
		return errors.WithContext("error writing to file: ", err)
	} else if err = f.Close(); err != nil {
		return errors.WithContext("error closing file: ", err)
	}
	if fs.tmpDir != "" {
		fp := f.Name()
		if err = os.Rename(fp, filepath.Join(fs.baseDir, key)); err != nil {
			os.Remove(fp)
			return errors.WithContext("error moving tmp file: ", err)
		}
	}
	return nil
}

func (fs *fileStore) Remove(key string) error {
	if err := testKey(key); err != nil {
		return err
	}
	if os.IsNotExist((os.Remove(filepath.Join(fs.baseDir, key)))) {
		return ErrUnknownKey
	}
	return nil
}

func testKey(key string) error {
	if strings.ContainsRune(key, filepath.Separator) {
		return ErrInvalidKey
	}
	return nil
}
