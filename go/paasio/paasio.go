package paasio

import (
	"io"
	"sync"
)

type readCounter struct {
	reader         io.Reader
	totalBytesRead int
	numReads       int
	mutex          *sync.RWMutex
}

// type writeCounter struct {
// 	writer io.Writer
// }

// type readWriteCounter struct {
// 	ReadCounter
// 	WriteCounter
// }

// func NewWriteCounter(writer io.Writer) WriteCounter {
// 	panic("Please implement the NewWriterCounter function")
// }

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{
		reader: reader,
		mutex:  new(sync.RWMutex),
	}
}

// func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
// 	return &readWriteCounter{
// 		NewReadCounter(readwriter),
// 		NewWriteCounter(readwriter),
// 	}
// }

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
	return int64(rc.totalBytesRead), rc.numReads
}

// func (wc *writeCounter) Write(p []byte) (int, error) {
// 	panic("Please implement the Write function")
// }

// func (wc *writeCounter) WriteCount() (int64, int) {
// 	panic("Please implement the WriteCount function")
// }
