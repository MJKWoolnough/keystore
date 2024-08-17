package keystore

// File automatically generated with ./types.sh.

import "io"

type value interface {
	io.ReaderFrom
	io.WriterTo
}

var (
	_ value = new(String)
	_ value = new(Uint8)
	_ value = new(Uint16)
	_ value = new(Uint32)
	_ value = new(Uint64)
	_ value = new(Uint)
	_ value = new(Int8)
	_ value = new(Int16)
	_ value = new(Int32)
	_ value = new(Int64)
	_ value = new(Int)
	_ value = new(Float32)
	_ value = new(Float64)
)
