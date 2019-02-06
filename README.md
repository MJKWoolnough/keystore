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

FileBackedMemStore combines both a FileStore and a MemStore

#### func  NewFileBackedMemStore

```go
func NewFileBackedMemStore(baseDir, tmpDir string, mangler Mangler) (*FileBackedMemStore, error)
```
NewFileBackedMemStore create a new Store which uses the filesystem for permanent
storage, but uses memory for caching

#### func (*FileBackedMemStore) Get

```go
func (fs *FileBackedMemStore) Get(key string, r io.ReaderFrom) error
```
Get retrieves a key from the Store, first looking in the memcache and then going
to the filesystem

#### func (*FileBackedMemStore) Keys

```go
func (fs *FileBackedMemStore) Keys() []string
```
Keys returns a sorted slice of all of the keys

#### func (*FileBackedMemStore) Remove

```go
func (fs *FileBackedMemStore) Remove(key string) error
```
Remove deletes a key from both the memcache and the filesystem

#### func (*FileBackedMemStore) Set

```go
func (fs *FileBackedMemStore) Set(key string, w io.WriterTo) error
```
Set stores the key in both the memcache and the filesystem

#### type FileStore

```go
type FileStore struct {
}
```

FileStore implements the Store interface and provides a file backed keystore

#### func  NewFileStore

```go
func NewFileStore(baseDir, tmpDir string, mangler Mangler) (*FileStore, error)
```
NewFileStore creates a file backed key-value store

#### func (*FileStore) Get

```go
func (fs *FileStore) Get(key string, r io.ReaderFrom) error
```
Get retrieves the key data from the filesystem

#### func (*FileStore) Keys

```go
func (fs *FileStore) Keys() []string
```
Keys returns a sorted slice of all of the keys

#### func (*FileStore) Remove

```go
func (fs *FileStore) Remove(key string) error
```
Remove deletes the key data from the filesystem

#### func (*FileStore) Set

```go
func (fs *FileStore) Set(key string, w io.WriterTo) error
```
Set stores the key data on the filesystem

#### type Float32

```go
type Float32 float32
```

Float32 is a float32 that implements io.ReaderFrom and io.WriterTo

#### func (*Float32) ReadFrom

```go
func (t *Float32) ReadFrom(r io.Reader) (int64, error)
```
ReadFrom decodes the float32 from the Reader

#### func (Float32) WriteTo

```go
func (t Float32) WriteTo(w io.Writer) (int64, error)
```
WriteTo encodes the float32 to the Writer

#### type Float64

```go
type Float64 float64
```

Float64 is a float64 that implements io.ReaderFrom and io.WriterTo

#### func (*Float64) ReadFrom

```go
func (t *Float64) ReadFrom(r io.Reader) (int64, error)
```
ReadFrom decodes the float64 from the Reader

#### func (Float64) WriteTo

```go
func (t Float64) WriteTo(w io.Writer) (int64, error)
```
WriteTo encodes the float64 to the Writer

#### type Int

```go
type Int int
```

Int is a int that implements io.ReaderFrom and io.WriterTo

#### func (*Int) ReadFrom

```go
func (t *Int) ReadFrom(r io.Reader) (int64, error)
```
ReadFrom decodes the int from the Reader

#### func (Int) WriteTo

```go
func (t Int) WriteTo(w io.Writer) (int64, error)
```
WriteTo encodes the int to the Writer

#### type Int16

```go
type Int16 int16
```

Int16 is a int16 that implements io.ReaderFrom and io.WriterTo

#### func (*Int16) ReadFrom

```go
func (t *Int16) ReadFrom(r io.Reader) (int64, error)
```
ReadFrom decodes the int16 from the Reader

#### func (Int16) WriteTo

```go
func (t Int16) WriteTo(w io.Writer) (int64, error)
```
WriteTo encodes the int16 to the Writer

#### type Int32

```go
type Int32 int32
```

Int32 is a int32 that implements io.ReaderFrom and io.WriterTo

#### func (*Int32) ReadFrom

```go
func (t *Int32) ReadFrom(r io.Reader) (int64, error)
```
ReadFrom decodes the int32 from the Reader

#### func (Int32) WriteTo

```go
func (t Int32) WriteTo(w io.Writer) (int64, error)
```
WriteTo encodes the int32 to the Writer

#### type Int64

```go
type Int64 int64
```

Int64 is a int64 that implements io.ReaderFrom and io.WriterTo

#### func (*Int64) ReadFrom

```go
func (t *Int64) ReadFrom(r io.Reader) (int64, error)
```
ReadFrom decodes the int64 from the Reader

#### func (Int64) WriteTo

```go
func (t Int64) WriteTo(w io.Writer) (int64, error)
```
WriteTo encodes the int64 to the Writer

#### type Int8

```go
type Int8 int8
```

Int8 is a int8 that implements io.ReaderFrom and io.WriterTo

#### func (*Int8) ReadFrom

```go
func (t *Int8) ReadFrom(r io.Reader) (int64, error)
```
ReadFrom decodes the int8 from the Reader

#### func (Int8) WriteTo

```go
func (t Int8) WriteTo(w io.Writer) (int64, error)
```
WriteTo encodes the int8 to the Writer

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

```go
var NoMangle Mangler = noMangle{}
```
NoMangle is a mangler that performs no mangling. This should only be used when
you are certain that there are no filesystem special characters in the key name

#### type MemStore

```go
type MemStore struct {
}
```

MemStore implements Store and does so entirely in memory

#### func  NewMemStore

```go
func NewMemStore() *MemStore
```
NewMemStore creates a new memory-backed key-value store

#### func (*MemStore) Get

```go
func (ms *MemStore) Get(key string, r io.ReaderFrom) error
```
Get retrieves the key data from memory

#### func (*MemStore) GetAll

```go
func (ms *MemStore) GetAll(data map[string]io.ReaderFrom) error
```
GetAll retrieves data for all of the keys given. Useful to reduce locking. If
any of the keys do not exist no data will be read.

#### func (*MemStore) Keys

```go
func (ms *MemStore) Keys() []string
```
Keys returns a sorted slice of all of the keys

#### func (*MemStore) ReadFrom

```go
func (ms *MemStore) ReadFrom(r io.Reader) (int64, error)
```
ReadFrom implements the io.ReaderFrom interface allowing a MemStore to be be
retrieved in another Store

#### func (*MemStore) Remove

```go
func (ms *MemStore) Remove(key string) error
```
Remove deletes the key data from memory

#### func (*MemStore) Set

```go
func (ms *MemStore) Set(key string, w io.WriterTo) error
```
Set stores the key data in memory

#### func (*MemStore) SetAll

```go
func (ms *MemStore) SetAll(data map[string]io.WriterTo) error
```
SetAll set data for all of the keys given. Useful to reduce locking. Will return
the first error found, so may not set all data.

#### func (*MemStore) WriteTo

```go
func (ms *MemStore) WriteTo(w io.Writer) (int64, error)
```
WriteTo implements the io.WriterTo interface allowing a MemStore to be be stored
in another Store

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

#### type String

```go
type String string
```

String is a string that implements io.ReaderFrom and io.WriterTo

#### func (*String) ReadFrom

```go
func (t *String) ReadFrom(r io.Reader) (int64, error)
```
ReadFrom decodes the string from the Reader

#### func (String) WriteTo

```go
func (t String) WriteTo(w io.Writer) (int64, error)
```
WriteTo encodes the string to the Writer

#### type Uint

```go
type Uint uint
```

Uint is a uint that implements io.ReaderFrom and io.WriterTo

#### func (*Uint) ReadFrom

```go
func (t *Uint) ReadFrom(r io.Reader) (int64, error)
```
ReadFrom decodes the uint from the Reader

#### func (Uint) WriteTo

```go
func (t Uint) WriteTo(w io.Writer) (int64, error)
```
WriteTo encodes the uint to the Writer

#### type Uint16

```go
type Uint16 uint16
```

Uint16 is a uint16 that implements io.ReaderFrom and io.WriterTo

#### func (*Uint16) ReadFrom

```go
func (t *Uint16) ReadFrom(r io.Reader) (int64, error)
```
ReadFrom decodes the uint16 from the Reader

#### func (Uint16) WriteTo

```go
func (t Uint16) WriteTo(w io.Writer) (int64, error)
```
WriteTo encodes the uint16 to the Writer

#### type Uint32

```go
type Uint32 uint32
```

Uint32 is a uint32 that implements io.ReaderFrom and io.WriterTo

#### func (*Uint32) ReadFrom

```go
func (t *Uint32) ReadFrom(r io.Reader) (int64, error)
```
ReadFrom decodes the uint32 from the Reader

#### func (Uint32) WriteTo

```go
func (t Uint32) WriteTo(w io.Writer) (int64, error)
```
WriteTo encodes the uint32 to the Writer

#### type Uint64

```go
type Uint64 uint64
```

Uint64 is a uint64 that implements io.ReaderFrom and io.WriterTo

#### func (*Uint64) ReadFrom

```go
func (t *Uint64) ReadFrom(r io.Reader) (int64, error)
```
ReadFrom decodes the uint64 from the Reader

#### func (Uint64) WriteTo

```go
func (t Uint64) WriteTo(w io.Writer) (int64, error)
```
WriteTo encodes the uint64 to the Writer

#### type Uint8

```go
type Uint8 uint8
```

Uint8 is a uint8 that implements io.ReaderFrom and io.WriterTo

#### func (*Uint8) ReadFrom

```go
func (t *Uint8) ReadFrom(r io.Reader) (int64, error)
```
ReadFrom decodes the uint8 from the Reader

#### func (Uint8) WriteTo

```go
func (t Uint8) WriteTo(w io.Writer) (int64, error)
```
WriteTo encodes the uint8 to the Writer
