# keystore

[![CI](https://github.com/MJKWoolnough/keystore/actions/workflows/go-checks.yml/badge.svg)](https://github.com/MJKWoolnough/keystore/actions)
[![Go Reference](https://pkg.go.dev/badge/vimagination.zapto.org/keystore.svg)](https://pkg.go.dev/vimagination.zapto.org/keystore)

--
    import "vimagination.zapto.org/keystore"

Package keystore is a simple key-value storage system with file and memory backing.

## Highlights

 - In memory and file backed keystores.
 - A combined file-backed, in memory keystore.
 - Supports base64 key encoding for filesystem safe keys.
 - Simple CRUD-like API.

## Usage

```go
package main

import (
	"bytes"
	"fmt"
	"strings"

	"vimagination.zapto.org/keystore"
	"vimagination.zapto.org/memio"
)

func main() {
	store := keystore.NewMemStore()

	store.Set("myKey", strings.NewReader("VALUE"))
	store.Set("myKey2", strings.NewReader("FooBar"))

	var buf bytes.Buffer

	store.WriteTo(&buf)

	store = keystore.NewMemStore()

	store.ReadFrom(&buf)

	var mr memio.Buffer

	if err := store.Get("myKey", &mr); err != nil {
		fmt.Println("Error: myKey", err)
	} else {
		fmt.Println("myKey", string(mr))
	}

	mr = mr[:0]

	if err := store.Get("myKey2", &mr); err != nil {
		fmt.Println("Error: myKey2", err)
	} else {
		fmt.Println("myKey2", string(mr))
	}

	mr = mr[:0]

	if err := store.Get("unknownKey", &mr); err != nil {
		fmt.Println("Error: unknownKey", err)
	} else {
		fmt.Println("unknownKey", string(mr))
	}

	// Output:
	// myKey VALUE
	// myKey2 FooBar
	// Error: unknownKey key not found
}
```

## Documentation

Full API docs can be found at:

https://pkg.go.dev/vimagination.zapto.org/keystore
