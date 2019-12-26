// Package keystore is a simple key-value storage system with file and memory backing
package keystore // import "vimagination.zapto.org/keystore"

import (
	"errors"
	"io"
)

// Store represents the methods required for a Keystore
type Store interface {
	Get(string, io.ReaderFrom) error
	Set(string, io.WriterTo) error
	Remove(string) error
	Keys() []string
	Rename(string, string) error
}

// Errors
var (
	ErrUnknownKey = errors.New("key not found")
	ErrKeyExists  = errors.New("key already exists")
	ErrInvalidKey = errors.New("key contains invalid characters")
)
