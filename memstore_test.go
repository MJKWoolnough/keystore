package keystore

import (
	"testing"
)

func TestMemStore(t *testing.T) {
	testStore(t, NewMemStore())
}
