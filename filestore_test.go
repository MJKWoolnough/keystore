package keystore

import (
	"testing"
)

func TestFileStore(t *testing.T) {
	s, err := NewFileStore(t.TempDir(), "", nil)
	if err != nil {
		t.Errorf("received unexpected error creating FileStore: %s", err)
		return
	}
	testStore(t, s)
}

func TestFileStoreWithTmp(t *testing.T) {
	s, err := NewFileStore(t.TempDir(), t.TempDir(), nil)
	if err != nil {
		t.Errorf("received unexpected error creating FileStore: %s", err)
		return
	}
	testStore(t, s)
}
