package paasio

import (
	"io"
	"sync"
)

// === ReadCounter ===

type readCounter struct {
	reader         io.Reader
	totalBytesRead int
	numReads       int
	mutex          *sync.RWMutex
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{
		reader: reader,
		mutex:  new(sync.RWMutex),
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	rc.mutex.Lock()
	defer rc.mutex.Unlock()
	bytesRead, err := rc.reader.Read(p)
	if err != nil {
		return 0, err
	}

	rc.totalBytesRead += bytesRead
	rc.numReads++

	return bytesRead, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.mutex.RLock()
	defer rc.mutex.RUnlock()
	return int64(rc.totalBytesRead), rc.numReads
}

// === WriteCounter ===

type writeCounter struct {
	writer            io.Writer
	totalBytesWritten int
	numWrites         int
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{
		writer: writer,
	}
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	written, err := wc.writer.Write(p)
	if err != nil {
		return 0, err
	}

	wc.totalBytesWritten += written
	wc.numWrites++

	return written, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	panic("Please implement the WriteCount function")
}

// === ReadWriteCounter ===

// type readWriteCounter struct {
// 	ReadCounter
// 	WriteCounter
// }

// func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
// 	return &readWriteCounter{
// 		NewReadCounter(readwriter),
// 		NewWriteCounter(readwriter),
// 	}
// }
