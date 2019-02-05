// Package keystore is a simple key-value storage system with file and memory backing
package keystore // import "vimagination.zapto.org/keystore"

import (
	"io"

	"vimagination.zapto.org/errors"
)

// Store represents the methods required for a Keystore
type Store interface {
	Get(string, io.ReaderFrom) error
	Set(string, io.WriterTo) error
	Remove(string) error
}

// Errors
const (
	ErrUnknownKey errors.Error = "key not found"
	ErrInvalidKey errors.Error = "key contains invalid characters"
)
