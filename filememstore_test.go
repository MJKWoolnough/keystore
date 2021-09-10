package keystore

import (
	"testing"
)

func TestFileMemStore(t *testing.T) {
	s, err := NewFileBackedMemStore(t.TempDir(), "", nil)
	if err != nil {
		t.Errorf("received unexpected error creating FileStore: %s", err)
		return
	}
	testStore(t, s)
}

func TestFileMemStoreWithTmp(t *testing.T) {
	s, err := NewFileBackedMemStore(t.TempDir(), t.TempDir(), nil)
	if err != nil {
		t.Errorf("received unexpected error creating FileStore: %s", err)
		return
	}
	testStore(t, s)
}
