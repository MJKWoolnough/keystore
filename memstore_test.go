package keystore

import (
	"io"
	"reflect"
	"testing"

	"vimagination.zapto.org/memio"
)

func TestMemStore(t *testing.T) {
	testStore(t, NewMemStore())
}

func data(s string) io.WriterTo {
	b := memio.Buffer(s)
	return &b
}

func TestMemStoreReadWrite(t *testing.T) {
	m := NewMemStore()
	m.Set("key1", data("data1"))
	m.Set("key2", data("data2"))
	m.Set("abc123", data("abcdefghij"))
	m.Set("zxy987", data("1234567890"))
	m.Set("aMuchLongerKey", data("lotsAndLotsAndLotsOfData"))
	var buf memio.Buffer
	m.WriteTo(&buf)
	n := NewMemStore()
	n.ReadFrom(&buf)
	keys := m.Keys()
	if nKeys := n.Keys(); !reflect.DeepEqual(keys, nKeys) {
		t.Errorf("test 1: expecting keys %v, got %v", keys, nKeys)
		return
	}
	for d, key := range keys {
		buf = buf[:0]
		if err := n.Get(key, &buf); err != nil {
			t.Errorf("test %d: got unexpected error: %s", d+2, err)
			return
		}
		s := string(buf)
		buf = buf[:0]
		m.Get(key, &buf)
		if string(buf) != s {
			t.Errorf("test %d: expecting value %q, got %q", d+2, buf, s)
		}
	}
}
