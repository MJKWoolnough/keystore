package keystore

// File automatically generated with ./types.sh.

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
	aReaderPool = readerPool{
		pool: sync.Pool{
			New: func() interface{} {
				return new(byteio.StickyLittleEndianReader)
			},
		},
	}
	aWriterPool = writerPool{
		pool: sync.Pool{
			New: func() interface{} {
				return new(byteio.StickyLittleEndianWriter)
			},
		},
	}
)

// String is a string that implements io.ReaderFrom and io.WriterTo.
type String string

// ReadFrom decodes the string from the Reader.
func (t *String) ReadFrom(r io.Reader) (int64, error) {
	lr := aReaderPool.Get(r)
	*t = String(lr.ReadStringX())

	return aReaderPool.Put(lr)
}

// WriteTo encodes the string to the Writer.
func (t String) WriteTo(w io.Writer) (int64, error) {
	lw := aWriterPool.Get(w)

	lw.WriteStringX(string(t))

	return aWriterPool.Put(lw)
}

// Uint8 is a uint8 that implements io.ReaderFrom and io.WriterTo.
type Uint8 uint8

// ReadFrom decodes the uint8 from the Reader.
func (t *Uint8) ReadFrom(r io.Reader) (int64, error) {
	lr := aReaderPool.Get(r)
	*t = Uint8(lr.ReadUint8())

	return aReaderPool.Put(lr)
}

// WriteTo encodes the uint8 to the Writer.
func (t Uint8) WriteTo(w io.Writer) (int64, error) {
	lw := aWriterPool.Get(w)

	lw.WriteUint8(uint8(t))

	return aWriterPool.Put(lw)
}

// Uint16 is a uint16 that implements io.ReaderFrom and io.WriterTo.
type Uint16 uint16

// ReadFrom decodes the uint16 from the Reader.
func (t *Uint16) ReadFrom(r io.Reader) (int64, error) {
	lr := aReaderPool.Get(r)
	*t = Uint16(lr.ReadUint16())

	return aReaderPool.Put(lr)
}

// WriteTo encodes the uint16 to the Writer.
func (t Uint16) WriteTo(w io.Writer) (int64, error) {
	lw := aWriterPool.Get(w)

	lw.WriteUint16(uint16(t))

	return aWriterPool.Put(lw)
}

// Uint32 is a uint32 that implements io.ReaderFrom and io.WriterTo.
type Uint32 uint32

// ReadFrom decodes the uint32 from the Reader.
func (t *Uint32) ReadFrom(r io.Reader) (int64, error) {
	lr := aReaderPool.Get(r)
	*t = Uint32(lr.ReadUint32())

	return aReaderPool.Put(lr)
}

// WriteTo encodes the uint32 to the Writer.
func (t Uint32) WriteTo(w io.Writer) (int64, error) {
	lw := aWriterPool.Get(w)

	lw.WriteUint32(uint32(t))

	return aWriterPool.Put(lw)
}

// Uint64 is a uint64 that implements io.ReaderFrom and io.WriterTo.
type Uint64 uint64

// ReadFrom decodes the uint64 from the Reader.
func (t *Uint64) ReadFrom(r io.Reader) (int64, error) {
	lr := aReaderPool.Get(r)
	*t = Uint64(lr.ReadUint64())

	return aReaderPool.Put(lr)
}

// WriteTo encodes the uint64 to the Writer.
func (t Uint64) WriteTo(w io.Writer) (int64, error) {
	lw := aWriterPool.Get(w)

	lw.WriteUint64(uint64(t))

	return aWriterPool.Put(lw)
}

// Uint is a uint that implements io.ReaderFrom and io.WriterTo.
type Uint uint

// ReadFrom decodes the uint from the Reader.
func (t *Uint) ReadFrom(r io.Reader) (int64, error) {
	lr := aReaderPool.Get(r)
	*t = Uint(lr.ReadUintX())

	return aReaderPool.Put(lr)
}

// WriteTo encodes the uint to the Writer.
func (t Uint) WriteTo(w io.Writer) (int64, error) {
	lw := aWriterPool.Get(w)

	lw.WriteUintX(uint64(t))

	return aWriterPool.Put(lw)
}

// Int8 is a int8 that implements io.ReaderFrom and io.WriterTo.
type Int8 int8

// ReadFrom decodes the int8 from the Reader.
func (t *Int8) ReadFrom(r io.Reader) (int64, error) {
	lr := aReaderPool.Get(r)
	*t = Int8(lr.ReadInt8())

	return aReaderPool.Put(lr)
}

// WriteTo encodes the int8 to the Writer.
func (t Int8) WriteTo(w io.Writer) (int64, error) {
	lw := aWriterPool.Get(w)

	lw.WriteInt8(int8(t))

	return aWriterPool.Put(lw)
}

// Int16 is a int16 that implements io.ReaderFrom and io.WriterTo.
type Int16 int16

// ReadFrom decodes the int16 from the Reader.
func (t *Int16) ReadFrom(r io.Reader) (int64, error) {
	lr := aReaderPool.Get(r)
	*t = Int16(lr.ReadInt16())

	return aReaderPool.Put(lr)
}

// WriteTo encodes the int16 to the Writer.
func (t Int16) WriteTo(w io.Writer) (int64, error) {
	lw := aWriterPool.Get(w)

	lw.WriteInt16(int16(t))

	return aWriterPool.Put(lw)
}

// Int32 is a int32 that implements io.ReaderFrom and io.WriterTo.
type Int32 int32

// ReadFrom decodes the int32 from the Reader.
func (t *Int32) ReadFrom(r io.Reader) (int64, error) {
	lr := aReaderPool.Get(r)
	*t = Int32(lr.ReadInt32())

	return aReaderPool.Put(lr)
}

// WriteTo encodes the int32 to the Writer.
func (t Int32) WriteTo(w io.Writer) (int64, error) {
	lw := aWriterPool.Get(w)

	lw.WriteInt32(int32(t))

	return aWriterPool.Put(lw)
}

// Int64 is a int64 that implements io.ReaderFrom and io.WriterTo.
type Int64 int64

// ReadFrom decodes the int64 from the Reader.
func (t *Int64) ReadFrom(r io.Reader) (int64, error) {
	lr := aReaderPool.Get(r)
	*t = Int64(lr.ReadInt64())

	return aReaderPool.Put(lr)
}

// WriteTo encodes the int64 to the Writer.
func (t Int64) WriteTo(w io.Writer) (int64, error) {
	lw := aWriterPool.Get(w)

	lw.WriteInt64(int64(t))

	return aWriterPool.Put(lw)
}

// Int is a int that implements io.ReaderFrom and io.WriterTo.
type Int int

// ReadFrom decodes the int from the Reader.
func (t *Int) ReadFrom(r io.Reader) (int64, error) {
	lr := aReaderPool.Get(r)
	*t = Int(lr.ReadIntX())

	return aReaderPool.Put(lr)
}

// WriteTo encodes the int to the Writer.
func (t Int) WriteTo(w io.Writer) (int64, error) {
	lw := aWriterPool.Get(w)

	lw.WriteIntX(int64(t))

	return aWriterPool.Put(lw)
}

// Float32 is a float32 that implements io.ReaderFrom and io.WriterTo.
type Float32 float32

// ReadFrom decodes the float32 from the Reader.
func (t *Float32) ReadFrom(r io.Reader) (int64, error) {
	lr := aReaderPool.Get(r)
	*t = Float32(lr.ReadFloat32())

	return aReaderPool.Put(lr)
}

// WriteTo encodes the float32 to the Writer.
func (t Float32) WriteTo(w io.Writer) (int64, error) {
	lw := aWriterPool.Get(w)

	lw.WriteFloat32(float32(t))

	return aWriterPool.Put(lw)
}

// Float64 is a float64 that implements io.ReaderFrom and io.WriterTo.
type Float64 float64

// ReadFrom decodes the float64 from the Reader.
func (t *Float64) ReadFrom(r io.Reader) (int64, error) {
	lr := aReaderPool.Get(r)
	*t = Float64(lr.ReadFloat64())

	return aReaderPool.Put(lr)
}

// WriteTo encodes the float64 to the Writer.
func (t Float64) WriteTo(w io.Writer) (int64, error) {
	lw := aWriterPool.Get(w)

	lw.WriteFloat64(float64(t))

	return aWriterPool.Put(lw)
}
