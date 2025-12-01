package reader

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// LineReader reads a file line by line
type LineReader struct {
	r *bufio.Reader
	f *os.File
}

// New creates a new LineReader for the given path with passed buffer size
func New(path string, bufSize int) (*LineReader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &LineReader{
		r: bufio.NewReaderSize(f, bufSize),
		f: f,
	}, nil
}

// Close closes file
func (lr *LineReader) Close() error {
	if lr.f == nil {
		return nil
	}

	return lr.f.Close()
}

// Read returns a channel of string with trimmed IP
// Channel will be closed in case EOF is reached
func (lr *LineReader) Read() (<-chan string, <-chan error) {
	linesCh := make(chan string)
	errCh := make(chan error, 1)

	go func() {
		defer close(linesCh)
		for {
			line, err := lr.r.ReadString('\n')
			if len(line) > 0 {
				ln := strings.TrimSpace(line)
				linesCh <- ln
			}
			if err != nil {
				if err == io.EOF {
					fmt.Println("stop reading lines in file -> EOF")
					return
				}
				errCh <- err
				return
			}
		}
	}()
	return linesCh, errCh
}
