package paasio

import (
	"io"
)

type readCounter struct {
	reader io.Reader
}

type writeCounter struct {
	writer io.Writer
}

type readWriteCounter struct {
	ReadCounter
	WriteCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	panic("Please implement the NewWriterCounter function")
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{
		reader: reader,
	}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		NewReadCounter(readwriter),
		NewWriteCounter(readwriter),
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	return rc.reader.Read(p)
}

func (rc *readCounter) ReadCount() (int64, int) {
	panic("Please implement the ReadCount function")
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	panic("Please implement the Write function")
}

func (wc *writeCounter) WriteCount() (int64, int) {
	panic("Please implement the WriteCount function")
}
