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

// FileStore implements the Store interface and provides a file backed keystore
type FileStore struct {
	baseDir, tmpDir string
	mangler         Mangler
}

// NewFileStore creates a file backed key-value store
func NewFileStore(baseDir, tmpDir string, mangler Mangler) (*FileStore, error) {
	fs := new(FileStore)
	if err := fs.init(baseDir, tmpDir, mangler); err != nil {
		return nil, err
	}
	return fs, nil
}

func (fs *FileStore) init(baseDir, tmpDir string, mangler Mangler) error {
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

// Get retrieves the key data from the filesystem
func (fs *FileStore) Get(key string, r io.ReaderFrom) error {
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

// Set stores the key data on the filesystem
func (fs *FileStore) Set(key string, w io.WriterTo) error {
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

// Remove deletes the key data from the filesystem
func (fs *FileStore) Remove(key string) error {
	key = fs.mangleKey(key, false)
	if os.IsNotExist((os.Remove(filepath.Join(fs.baseDir, key)))) {
		return ErrUnknownKey
	}
	return nil
}

// Keys returns a sorted slice of all of the keys
func (fs *FileStore) Keys() []string {
	s := fs.getDirContents("")
	sort.Strings(s)
	return s
}

// Stat returns the FileInfo of the file relatining to the given key
func (fs *FileStore) Stat(key string) (os.FileInfo, error) {
	return os.Stat(filepath.Join(fs.baseDir, fs.mangleKey(key, false)))
}

// Exists returns true when the key exists within the store
func (fs *FileStore) Exists(key string) bool {
	_, err := os.Stat(filepath.Join(fs.baseDir, fs.mangleKey(key, false)))
	return err == nil
}

func (fs *FileStore) mangleKey(key string, prepare bool) string {
	parts := fs.mangler.Encode(key)
	if len(parts) == 0 {
		return ""
	} else if len(parts) == 1 {
		return parts[0]
	} else if prepare {
		os.MkdirAll(filepath.Join(append([]string{fs.baseDir}, parts...)...), 0700)
	}
	return filepath.Clean("/" + strings.Join(parts, string(filepath.Separator)))[1:]
}

func (fs *FileStore) getDirContents(dir string) []string {
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

// Mangler is an interface for the methods required to un/mangle a key
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

type noMangle struct{}

func (noMangle) Encode(name string) []string {
	return strings.Split(name, string(filepath.Separator))
}

func (noMangle) Decode(parts []string) (string, error) {
	return strings.Join(parts, string(filepath.Separator)), nil
}

// Base64Mangler represents the default Mangler that simple base64 encodes the
// key
var Base64Mangler Mangler = base64Mangler{}

// NoMangle is a mangler that performs no mangling. This should only be used
// when you are certain that there are no filesystem special characters in the
// key name
var NoMangle Mangler = noMangle{}
