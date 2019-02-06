package keystore

// File automatically generated with ./types.sh

import (
	"io"

	"vimagination.zapto.org/byteio"
)

// String is a string that implements io.ReaderFrom and io.WriterTo
type String string

// ReadFrom decodes the string from the Reader
func (t *String) ReadFrom(r io.Reader) (int64, error) {
	lr := byteio.StickyLittleEndianReader{Reader: r}
	*t = String(lr.ReadStringX())
	return lr.Count, lr.Err
}

// WriteTo encodes the string to the Writer
func (t String) WriteTo(w io.Writer) (int64, error) {
	lw := byteio.StickyLittleEndianWriter{Writer: w}
	lw.WriteStringX(string(t))
	return lw.Count, lw.Err
}

// Uint8 is a uint8 that implements io.ReaderFrom and io.WriterTo
type Uint8 uint8

// ReadFrom decodes the uint8 from the Reader
func (t *Uint8) ReadFrom(r io.Reader) (int64, error) {
	lr := byteio.StickyLittleEndianReader{Reader: r}
	*t = Uint8(lr.ReadUint8())
	return lr.Count, lr.Err
}

// WriteTo encodes the uint8 to the Writer
func (t Uint8) WriteTo(w io.Writer) (int64, error) {
	lw := byteio.StickyLittleEndianWriter{Writer: w}
	lw.WriteUint8(uint8(t))
	return lw.Count, lw.Err
}

// Uint16 is a uint16 that implements io.ReaderFrom and io.WriterTo
type Uint16 uint16

// ReadFrom decodes the uint16 from the Reader
func (t *Uint16) ReadFrom(r io.Reader) (int64, error) {
	lr := byteio.StickyLittleEndianReader{Reader: r}
	*t = Uint16(lr.ReadUint16())
	return lr.Count, lr.Err
}

// WriteTo encodes the uint16 to the Writer
func (t Uint16) WriteTo(w io.Writer) (int64, error) {
	lw := byteio.StickyLittleEndianWriter{Writer: w}
	lw.WriteUint16(uint16(t))
	return lw.Count, lw.Err
}

// Uint32 is a uint32 that implements io.ReaderFrom and io.WriterTo
type Uint32 uint32

// ReadFrom decodes the uint32 from the Reader
func (t *Uint32) ReadFrom(r io.Reader) (int64, error) {
	lr := byteio.StickyLittleEndianReader{Reader: r}
	*t = Uint32(lr.ReadUint32())
	return lr.Count, lr.Err
}

// WriteTo encodes the uint32 to the Writer
func (t Uint32) WriteTo(w io.Writer) (int64, error) {
	lw := byteio.StickyLittleEndianWriter{Writer: w}
	lw.WriteUint32(uint32(t))
	return lw.Count, lw.Err
}

// Uint64 is a uint64 that implements io.ReaderFrom and io.WriterTo
type Uint64 uint64

// ReadFrom decodes the uint64 from the Reader
func (t *Uint64) ReadFrom(r io.Reader) (int64, error) {
	lr := byteio.StickyLittleEndianReader{Reader: r}
	*t = Uint64(lr.ReadUint64())
	return lr.Count, lr.Err
}

// WriteTo encodes the uint64 to the Writer
func (t Uint64) WriteTo(w io.Writer) (int64, error) {
	lw := byteio.StickyLittleEndianWriter{Writer: w}
	lw.WriteUint64(uint64(t))
	return lw.Count, lw.Err
}

// Uint is a uint that implements io.ReaderFrom and io.WriterTo
type Uint uint

// ReadFrom decodes the uint from the Reader
func (t *Uint) ReadFrom(r io.Reader) (int64, error) {
	lr := byteio.StickyLittleEndianReader{Reader: r}
	*t = Uint(lr.ReadUintX())
	return lr.Count, lr.Err
}

// WriteTo encodes the uint to the Writer
func (t Uint) WriteTo(w io.Writer) (int64, error) {
	lw := byteio.StickyLittleEndianWriter{Writer: w}
	lw.WriteUintX(uint64(t))
	return lw.Count, lw.Err
}

// Int8 is a int8 that implements io.ReaderFrom and io.WriterTo
type Int8 int8

// ReadFrom decodes the int8 from the Reader
func (t *Int8) ReadFrom(r io.Reader) (int64, error) {
	lr := byteio.StickyLittleEndianReader{Reader: r}
	*t = Int8(lr.ReadInt8())
	return lr.Count, lr.Err
}

// WriteTo encodes the int8 to the Writer
func (t Int8) WriteTo(w io.Writer) (int64, error) {
	lw := byteio.StickyLittleEndianWriter{Writer: w}
	lw.WriteInt8(int8(t))
	return lw.Count, lw.Err
}

// Int16 is a int16 that implements io.ReaderFrom and io.WriterTo
type Int16 int16

// ReadFrom decodes the int16 from the Reader
func (t *Int16) ReadFrom(r io.Reader) (int64, error) {
	lr := byteio.StickyLittleEndianReader{Reader: r}
	*t = Int16(lr.ReadInt16())
	return lr.Count, lr.Err
}

// WriteTo encodes the int16 to the Writer
func (t Int16) WriteTo(w io.Writer) (int64, error) {
	lw := byteio.StickyLittleEndianWriter{Writer: w}
	lw.WriteInt16(int16(t))
	return lw.Count, lw.Err
}

// Int32 is a int32 that implements io.ReaderFrom and io.WriterTo
type Int32 int32

// ReadFrom decodes the int32 from the Reader
func (t *Int32) ReadFrom(r io.Reader) (int64, error) {
	lr := byteio.StickyLittleEndianReader{Reader: r}
	*t = Int32(lr.ReadInt32())
	return lr.Count, lr.Err
}

// WriteTo encodes the int32 to the Writer
func (t Int32) WriteTo(w io.Writer) (int64, error) {
	lw := byteio.StickyLittleEndianWriter{Writer: w}
	lw.WriteInt32(int32(t))
	return lw.Count, lw.Err
}

// Int64 is a int64 that implements io.ReaderFrom and io.WriterTo
type Int64 int64

// ReadFrom decodes the int64 from the Reader
func (t *Int64) ReadFrom(r io.Reader) (int64, error) {
	lr := byteio.StickyLittleEndianReader{Reader: r}
	*t = Int64(lr.ReadInt64())
	return lr.Count, lr.Err
}

// WriteTo encodes the int64 to the Writer
func (t Int64) WriteTo(w io.Writer) (int64, error) {
	lw := byteio.StickyLittleEndianWriter{Writer: w}
	lw.WriteInt64(int64(t))
	return lw.Count, lw.Err
}

// Int is a int that implements io.ReaderFrom and io.WriterTo
type Int int

// ReadFrom decodes the int from the Reader
func (t *Int) ReadFrom(r io.Reader) (int64, error) {
	lr := byteio.StickyLittleEndianReader{Reader: r}
	*t = Int(lr.ReadIntX())
	return lr.Count, lr.Err
}

// WriteTo encodes the int to the Writer
func (t Int) WriteTo(w io.Writer) (int64, error) {
	lw := byteio.StickyLittleEndianWriter{Writer: w}
	lw.WriteIntX(int64(t))
	return lw.Count, lw.Err
}

// Float32 is a float32 that implements io.ReaderFrom and io.WriterTo
type Float32 float32

// ReadFrom decodes the float32 from the Reader
func (t *Float32) ReadFrom(r io.Reader) (int64, error) {
	lr := byteio.StickyLittleEndianReader{Reader: r}
	*t = Float32(lr.ReadFloat32())
	return lr.Count, lr.Err
}

// WriteTo encodes the float32 to the Writer
func (t Float32) WriteTo(w io.Writer) (int64, error) {
	lw := byteio.StickyLittleEndianWriter{Writer: w}
	lw.WriteFloat32(float32(t))
	return lw.Count, lw.Err
}

// Float64 is a float64 that implements io.ReaderFrom and io.WriterTo
type Float64 float64

// ReadFrom decodes the float64 from the Reader
func (t *Float64) ReadFrom(r io.Reader) (int64, error) {
	lr := byteio.StickyLittleEndianReader{Reader: r}
	*t = Float64(lr.ReadFloat64())
	return lr.Count, lr.Err
}

// WriteTo encodes the float64 to the Writer
func (t Float64) WriteTo(w io.Writer) (int64, error) {
	lw := byteio.StickyLittleEndianWriter{Writer: w}
	lw.WriteFloat64(float64(t))
	return lw.Count, lw.Err
}
