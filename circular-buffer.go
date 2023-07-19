package exercism

import (
	"errors"
)

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Define the Buffer type here.
type Buffer struct {
	Data       []byte
	Count      int
	Size       int
	WritePoint int
	ReadPoint  int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{
		Data:       make([]byte, size),
		Count:      0,
		Size:       size,
		WritePoint: 0,
		ReadPoint:  0,
	}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.Count == 0 {
		return 0, errors.New("The buffer is empty")
	}

	data := b.Data[b.ReadPoint]
	b.Count--
	b.ReadPoint = (b.ReadPoint + 1) % b.Size
	return data, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.Count == b.Size {
		return errors.New("The buffer is full")
	}

	b.Data[b.WritePoint] = c
	b.Count++
	b.WritePoint = (b.WritePoint + 1) % b.Size
	return nil
}

func (b *Buffer) Overwrite(c byte) {

	if b.Count < b.Size {
		b.WriteByte(c)
	} else {
		b.Data[b.ReadPoint] = c
		b.WritePoint = (b.WritePoint + 1) % b.Size
		b.ReadPoint = (b.ReadPoint + 1) % b.Size
	}
}

func (b *Buffer) Reset() {
	b.Count = 0
	b.ReadPoint = 0
	b.WritePoint = 0
}
