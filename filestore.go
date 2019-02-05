package keystore

import (
	"encoding/base64"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"vimagination.zapto.org/errors"
)

type fileStore struct {
	baseDir, tmpDir string
	mangler         Mangler
}

// NewFileStore creates a file backed key-value store
func NewFileStore(baseDir, tmpDir string, mangler Mangler) (Store, error) {
	fs := new(fileStore)
	if err := fs.init(baseDir, tmpDir, mangler); err != nil {
		return nil, err
	}
	return fs, nil
}

func (fs *fileStore) init(baseDir, tmpDir string, mangler Mangler) error {
	if err := os.MkdirAll(baseDir, 0700); err != nil {
		return errors.WithContext("error creating data dir: ", err)
	}
	if mangler == nil {
		mangler = base64Mangler{}
	}
	if tmpDir != "" {
		if err := os.MkdirAll(tmpDir, 0700); err != nil {
			return errors.WithContext("error creating temp dir: ", err)
		}
	}
	fs.baseDir = baseDir
	fs.tmpDir = tmpDir
	fs.mangler = mangler
	return nil
}

func (fs *fileStore) Get(key string, r io.ReaderFrom) error {
	key = fs.mangleKey(key, false)
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
	key = fs.mangleKey(key, true)
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
	key = fs.mangleKey(key, false)
	if os.IsNotExist((os.Remove(filepath.Join(fs.baseDir, key)))) {
		return ErrUnknownKey
	}
	return nil
}

// Keys returns a sorted slice of all of the keys
func (fs *fileStore) Keys() []string {
	s := fs.getDirContents("")
	sort.Strings(s)
	return s
}

func (fs *fileStore) mangleKey(key string, prepare bool) string {
	parts := fs.mangler.Encode(key)
	if len(parts) == 0 {
		return ""
	} else if len(parts) == 1 {
		return parts[0]
	} else if prepare {
		os.MkdirAll(filepath.Join(append([]string{fs.baseDir}, parts...)...), 0700)
	}
	return strings.Join(parts, string(filepath.Separator))
}

func (fs *fileStore) getDirContents(dir string) []string {
	d, err := os.Open(filepath.Join(fs.baseDir, dir))
	if err != nil {
		return nil
	}
	files, err := d.Readdir(-1)
	if err != nil {
		return nil
	}
	names := make([]string, 0, len(files))
	for _, file := range files {
		if file.IsDir() {
			names = append(names, fs.getDirContents(filepath.Join(dir, file.Name()))...)
		} else {
			name, err := fs.mangler.Decode(strings.Split(filepath.Join(dir, file.Name()), string(filepath.Separator)))
			if err != nil {
				continue
			}
			names = append(names, name)
		}
	}
	return names
}

type Mangler interface {
	Encode(string) []string
	Decode([]string) (string, error)
}

type base64Mangler struct{}

func (base64Mangler) Encode(name string) []string {
	return []string{base64.URLEncoding.EncodeToString([]byte(name))}
}

func (base64Mangler) Decode(parts []string) (string, error) {
	if len(parts) != 1 {
		return "", ErrInvalidKey
	}
	b, err := base64.URLEncoding.DecodeString(parts[0])
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Default Base64 mangler
var Base64Mangler Mangler = base64Mangler{}
