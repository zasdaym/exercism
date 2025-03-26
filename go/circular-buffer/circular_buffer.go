package circular

import "fmt"

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Define the Buffer type here.
type Buffer struct {
	data []byte
	head int
	tail int
	size int
}

func NewBuffer(size int) *Buffer {
	data := make([]byte, size)
	return &Buffer{data: data}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.size == 0 {
		return 0, fmt.Errorf("buffer is empty")
	}
	b.size -= 1
	c := b.data[b.tail]
	b.tail = (b.tail + 1) % len(b.data)
	return c, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.isFull() {
		return fmt.Errorf("buffer is full")
	}
	b.size += 1
	b.data[b.head] = c
	b.head = (b.head + 1) % len(b.data)
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if !b.isFull() {
		b.WriteByte(c)
		return
	}
	b.data[b.head] = c
	b.tail = (b.tail + 1) % len(b.data)
	b.head = (b.head + 1) % len(b.data)
}

func (b *Buffer) Reset() {
	b.head = 0
	b.tail = 0
	b.size = 0
}

func (b *Buffer) isFull() bool {
	return b.size == len(b.data)
}
