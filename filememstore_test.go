package keystore

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileMemStore(t *testing.T) {
	dir, err := ioutil.TempDir("", "fileStore-test")
	if err != nil {
		t.Errorf("received unexpected error creating temp dir: %s", err)
		return
	}
	defer os.RemoveAll(dir)
	s, err := NewFileBackedMemStore(dir, "")
	if err != nil {
		t.Errorf("received unexpected error creating FileStore: %s", err)
		return
	}
	testStore(t, s)
}

func TestFileMemStoreWithTmp(t *testing.T) {
	dir, err := ioutil.TempDir("", "fileStore-test")
	if err != nil {
		t.Errorf("received unexpected error creating temp dir: %s", err)
		return
	}
	defer os.RemoveAll(dir)
	tmp, err := ioutil.TempDir("", "fileStore-test-tmp")
	if err != nil {
		t.Errorf("received unexpected error creating tmp dir: %s", err)
		return
	}
	defer os.RemoveAll(tmp)
	s, err := NewFileBackedMemStore(dir, "")
	if err != nil {
		t.Errorf("received unexpected error creating FileStore: %s", err)
		return
	}
	testStore(t, s)
}
