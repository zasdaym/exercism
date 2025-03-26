package paasio

import (
	"io"
	"sync"
)

type readCounter struct {
	mu             sync.Mutex
	r              io.Reader
	readTotalBytes int64
	readTotalOps   int
}

type writeCounter struct {
	mu              sync.Mutex
	w               io.Writer
	writeTotalBytes int64
	writeTotalOps   int
}

type readWriteCounter struct {
	readCounter
	writeCounter
}

func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{w: w}
}

func NewReadCounter(r io.Reader) ReadCounter {
	return &readCounter{r: r}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		readCounter:  readCounter{r: readwriter},
		writeCounter: writeCounter{w: readwriter},
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	n, err := rc.r.Read(p)
	rc.readTotalBytes += int64(n)
	rc.readTotalOps++

	return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	return rc.readTotalBytes, rc.readTotalOps
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	wc.mu.Lock()
	defer wc.mu.Unlock()

	n, err := wc.w.Write(p)
	wc.writeTotalBytes += int64(n)
	wc.writeTotalOps++

	return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.mu.Lock()
	defer wc.mu.Unlock()

	return wc.writeTotalBytes, wc.writeTotalOps
}
