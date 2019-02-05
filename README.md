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

#### func  NewFileBackedMemStore

```go
func NewFileBackedMemStore(baseDir, tmpDir string) (Store, error)
```
NewFileBackedMemStore create s new Store which uses the filesystem for permanent
storage, but uses memory for caching

#### func  NewFileStore

```go
func NewFileStore(baseDir, tmpDir string) (Store, error)
```
NewFileStore creates a file backed key-value store

#### func  NewMemStore

```go
func NewMemStore() Store
```
NewMemStore creates a new memory-backed key-value store
