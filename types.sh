#!/bin/bash

allTypes="String Uint8 Uint16 Uint32 Uint64 Int8 Int16 Int32 Int64 Float32 Float64";

(
	cat <<HEREDOC
package keystore

// File automatically generated with ./types.sh

import (
	"io"

	"vimagination.zapto.org/byteio"
)
HEREDOC
	for typeName in $allTypes; do
		type="$(echo -n "$typeName" | tr A-Z a-z)";
		fName="$typeName";
		[ "$type" = "string" ] && fName="${typeName}X";
		cat <<HEREDOC

// $typeName is a $type that implements io.ReaderFrom and io.WriterTo
type $typeName $type

// ReadFrom decodes the $type from the Reader
func (t *$typeName) ReadFrom(r io.Reader) (int64, error) {
	lr := byteio.StickyLittleEndianReader{Reader: r}
	*t = $typeName(lr.Read$fName())
	return lr.Count, lr.Err
}

// WriteTo encodes the $type to the Writer
func (t $typeName) WriteTo(w io.Writer) (int64, error) {
	lw := byteio.StickyLittleEndianWriter{Writer: w}
	lw.Write$fName($type(t))
	return lw.Count, lw.Err
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
