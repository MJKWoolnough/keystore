#!/bin/bash

allTypes="String Uint8 Uint16 Uint32 Uint64 Uint Int8 Int16 Int32 Int64 Int Float32 Float64";

(
	cat <<HEREDOC
package keystore

// File automatically generated with ./types.sh

import (
	"io"
	"sync"

	"vimagination.zapto.org/byteio"
)

type readerPool struct {
	pool sync.Pool
}

func (rp *readerPool) Get(r io.Reader) *byteio.StickyLittleEndianReader {
	lr := rp.pool.Get().(*byteio.StickyLittleEndianReader)
	lr.Reader = r
	return lr
}

func (rp *readerPool) Put(lr *byteio.StickyLittleEndianReader) (int64, error) {
	c, err := lr.Count, lr.Err
	*lr = byteio.StickyLittleEndianReader{}
	rp.pool.Put(lr)
	return c, err
}

type writerPool struct {
	pool sync.Pool
}

func (wp *writerPool) Get(w io.Writer) *byteio.StickyLittleEndianWriter {
	lw := wp.pool.Get().(*byteio.StickyLittleEndianWriter)
	lw.Writer = w
	return lw
}

func (wp *writerPool) Put(lw *byteio.StickyLittleEndianWriter) (int64, error) {
	c, err := lw.Count, lw.Err
	*lw = byteio.StickyLittleEndianWriter{}
	wp.pool.Put(lw)
	return c, err
}

var (
	aReaderPool = readerPool {
		pool: sync.Pool {
			New: func() interface{} {
				return new(byteio.StickyLittleEndianReader)
			},
		},
	}
	aWriterPool = writerPool {
		pool: sync.Pool {
			New: func() interface {} {
				return new(byteio.StickyLittleEndianWriter)
			},
		},
	}
)
HEREDOC
	for typeName in $allTypes; do
		type="$(echo -n "$typeName" | tr A-Z a-z)";
		fName="$typeName";
		[ "$type" = "string" ] && fName="${typeName}X";
		wType="$type";
		[ "$type" = "uint" -o "$type" = "int" ] && {
			fName="${typeName}X";
			wType="${type}64";
		}
		cat <<HEREDOC

// $typeName is a $type that implements io.ReaderFrom and io.WriterTo
type $typeName $type

// ReadFrom decodes the $type from the Reader
func (t *$typeName) ReadFrom(r io.Reader) (int64, error) {
	lr := aReaderPool.Get(r)
	*t = $typeName(lr.Read$fName())
	return aReaderPool.Put(lr)
}

// WriteTo encodes the $type to the Writer
func (t $typeName) WriteTo(w io.Writer) (int64, error) {
	lw := aWriterPool.Get(w)
	lw.Write$fName($wType(t))
	return aWriterPool.Put(lw)
}
HEREDOC
	done;
) > types.go

(
	cat <<HEREDOC
package keystore

// File automatically generated with ./types.sh

import "io"

type value interface {
	io.ReaderFrom
	io.WriterTo
}

var (
HEREDOC
	for typeName in $allTypes; do
		echo "	_ value = new($typeName)";
	done;
	echo ")";
) > types_test.go;
