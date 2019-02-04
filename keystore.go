package keystore

import (
	"io"

	"vimagination.zapto.org/errors"
)

type Store interface {
	Get(string, io.ReaderFrom) error
	Set(string, io.WriterTo) error
	Remove(string) error
}

const (
	ErrUnknownKey errors.Error = "key not found"
	ErrInvalidKey errors.Error = "key contains invalid characters"
)
