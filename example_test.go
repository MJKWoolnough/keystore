package keystore_test

import (
	"bytes"
	"fmt"
	"strings"

	"vimagination.zapto.org/keystore"
	"vimagination.zapto.org/memio"
)

func Example() {
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
