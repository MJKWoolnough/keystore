package keystore

import (
	"reflect"
	"testing"

	"vimagination.zapto.org/memio"
)

func testStore(t *testing.T, s Store) {
	const testData = "Hello, World!"
	var buf memio.Buffer
	if err := s.Get("none", &buf); err != ErrUnknownKey {
		t.Errorf("test 1: expecting error ErrUnknownKey, got %s", err)
	} else if len(buf) > 0 {
		t.Errorf("test 1: received data when expecting none.")
	} else if err = s.Remove("none"); err != ErrUnknownKey {
		t.Errorf("test 2: expecting error ErrUnknownKey, got %s", err)
	} else if err = s.Set("key1", &buf); err != nil {
		t.Errorf("test 3: unexpected error: %s", err)
	} else if ss := s.Keys(); !reflect.DeepEqual(ss, []string{"key1"}) {
		t.Errorf("test 3: expecting key in Store")
	} else if err = s.Get("key1", &buf); err != nil {
		t.Errorf("test 4: unexpected error: %s", err)
	} else if len(buf) > 0 {
		t.Errorf("test 5: received data when expecting none.")
	} else {
		buf = memio.Buffer(testData)
		if err = s.Set("key2", &buf); err != nil {
			t.Errorf("test 6: unexpected error: %s", err)
		} else if len(buf) > 0 {
			t.Errorf("test 6: unread data: %v", err)
		} else if ss := s.Keys(); !reflect.DeepEqual(ss, []string{"key1", "key2"}) {
			t.Errorf("test 6: expecting key in Store")
		} else if err = s.Get("key2", &buf); err != nil {
			t.Errorf("test 7: unexpected error: %s", err)
		} else if string(buf) != testData {
			t.Errorf("test 7: expecting to read %q, got %v", testData, buf)
		} else if err = s.Remove("key2"); err != nil {
			t.Errorf("test 8: unexpected error: %s", err)
		} else if err = s.Get("key2", &buf); err != ErrUnknownKey {
			t.Errorf("test 8: key2 not removed: %s", err)
		}
	}
}
