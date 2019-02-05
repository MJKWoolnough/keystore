package keystore

import (
	"encoding/base64"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

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
	if err := os.MkdirAll(fs.baseDir, 0700); err != nil {
		return errors.WithContext("error creating data dir: ", err)
	}
	if tmpDir != "" {
		if err := os.MkdirAll(tmpDir, 0700); err != nil {
			return errors.WithContext("error creating temp dir: ", err)
		}
	}
	return nil
}

func (fs *fileStore) Get(key string, r io.ReaderFrom) error {
	key = base64.URLEncoding.EncodeToString([]byte(key))
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
	key = base64.URLEncoding.EncodeToString([]byte(key))
	var (
		f   *os.File
		err error
	)
	if fs.tmpDir != "" {
		f, err = ioutil.TempFile(fs.tmpDir, "keystore")
	} else {
		f, err = os.Create(filepath.Join(fs.baseDir, key))
	}
	if err != nil {
		return errors.WithContext("error opening file for writing: ", err)
	}
	if _, err = w.WriteTo(f); err != nil && err != io.EOF {
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
	key = base64.URLEncoding.EncodeToString([]byte(key))
	if os.IsNotExist((os.Remove(filepath.Join(fs.baseDir, key)))) {
		return ErrUnknownKey
	}
	return nil
}

// Keys returns a sorted slice of all of the keys
func (fs *fileStore) Keys() []string {
	d, err := os.Open(fs.baseDir)
	if err != nil {
		return nil
	}
	s, err := d.Readdirnames(-1)
	if err != nil {
		return nil
	}
	ss := make([]string, 0, len(s))
	for _, name := range s {
		bname, err := base64.URLEncoding.DecodeString(name)
		if err != nil {
			continue
		}
		ss = append(ss, string(bname))
	}
	sort.Strings(ss)
	return ss
}
