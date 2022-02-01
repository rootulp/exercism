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
	capacity    int
	store       []byte
	length      int // tracks the number of units in the buffer
	insertIndex int
	readIndex   int
}

func NewBuffer(capacity int) *Buffer {
	store := make([]byte, capacity)
	return &Buffer{capacity: capacity, store: store}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.len() == 0 {
		return 0, fmt.Errorf("error reading from empty buffer")
	}
	result := b.store[b.readIndex]
	b.store[b.readIndex] = 0
	b.readIndex = nextIndex(b.readIndex, b.capacity)
	b.length--
	return result, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.len() == b.capacity {
		return fmt.Errorf("error writing to full buffer")
	}
	b.store[b.insertIndex] = c
	b.insertIndex = nextIndex(b.insertIndex, b.capacity)
	b.length++
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if b.store[b.insertIndex] == 0 {
		b.length++
	} else {
		b.readIndex = nextIndex(b.readIndex, b.capacity)
	}
	b.store[b.insertIndex] = c
	b.insertIndex = nextIndex(b.insertIndex, b.capacity)
}

func (b *Buffer) Reset() {
	b.store = make([]byte, b.capacity)
	b.length = 0
	b.insertIndex = 0
	b.readIndex = 0
}

func (b *Buffer) len() int {
	return b.length
}

func nextIndex(current int, capacity int) int {
	// If we are at the last index in the buffer, loop around back to the 0
	// index
	if current == capacity-1 {
		return 0
	}
	return current + 1
}
