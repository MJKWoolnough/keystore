# keystore
--
    import "vimagination.zapto.org/keystore"

Package keystore is a simple key-value storage system with file and memory
### backing

## Usage

```go
const (
	ErrUnknownKey errors.Error = "key not found"
	ErrInvalidKey errors.Error = "key contains invalid characters"
)
```
Errors

#### type FileBackedMemStore

```go
type FileBackedMemStore struct {
}
```


#### func  NewFileBackedMemStore

```go
func NewFileBackedMemStore(baseDir, tmpDir string, mangler Mangler) (*FileBackedMemStore, error)
```
NewFileBackedMemStore create s new Store which uses the filesystem for permanent
storage, but uses memory for caching

#### func (*FileBackedMemStore) Get

```go
func (fs *FileBackedMemStore) Get(key string, r io.ReaderFrom) error
```

#### func (*FileBackedMemStore) Keys

```go
func (fs *FileBackedMemStore) Keys() []string
```
Keys returns a sorted slice of all of the keys

#### func (*FileBackedMemStore) Remove

```go
func (fs *FileBackedMemStore) Remove(key string) error
```

#### func (*FileBackedMemStore) Set

```go
func (fs *FileBackedMemStore) Set(key string, w io.WriterTo) error
```

#### type FileStore

```go
type FileStore struct {
}
```


#### func  NewFileStore

```go
func NewFileStore(baseDir, tmpDir string, mangler Mangler) (*FileStore, error)
```
NewFileStore creates a file backed key-value store

#### func (*FileStore) Get

```go
func (fs *FileStore) Get(key string, r io.ReaderFrom) error
```

#### func (*FileStore) Keys

```go
func (fs *FileStore) Keys() []string
```
Keys returns a sorted slice of all of the keys

#### func (*FileStore) Remove

```go
func (fs *FileStore) Remove(key string) error
```

#### func (*FileStore) Set

```go
func (fs *FileStore) Set(key string, w io.WriterTo) error
```

#### type Mangler

```go
type Mangler interface {
	Encode(string) []string
	Decode([]string) (string, error)
}
```

Mangler is an interface for the methods required to un/mangle a key

```go
var Base64Mangler Mangler = base64Mangler{}
```
Base64Mangler represents the default Mangler that simple base64 encodes the key

#### type MemStore

```go
type MemStore struct {
}
```


#### func  NewMemStore

```go
func NewMemStore() *MemStore
```
NewMemStore creates a new memory-backed key-value store

#### func (*MemStore) Get

```go
func (ms *MemStore) Get(key string, r io.ReaderFrom) error
```

#### func (*MemStore) Keys

```go
func (ms *MemStore) Keys() []string
```
Keys returns a sorted slice of all of the keys

#### func (*MemStore) ReadFrom

```go
func (ms *MemStore) ReadFrom(r io.Reader) (int64, error)
```

#### func (*MemStore) Remove

```go
func (ms *MemStore) Remove(key string) error
```

#### func (*MemStore) Set

```go
func (ms *MemStore) Set(key string, w io.WriterTo) error
```

#### func (*MemStore) WriteTo

```go
func (ms *MemStore) WriteTo(w io.Writer) (int64, error)
```

#### type Store

```go
type Store interface {
	Get(string, io.ReaderFrom) error
	Set(string, io.WriterTo) error
	Remove(string) error
	Keys() []string
}
```

Store represents the methods required for a Keystore
