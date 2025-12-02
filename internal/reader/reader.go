package reader

import (
	"bufio"
	"io"
	"os"
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

// Read is a wraper for ReadBytes with additional error checks
func (lr *LineReader) Read() ([]byte, error) {
	line, err := lr.r.ReadBytes('\n')
	if err != nil {
		if err == io.EOF && len(line) > 0 {
			return line, nil
		}
		return nil, err
	}
	return line, nil
}

// readSlices a wrapper for ReadSlice functionality with additional error handling
func (lr *LineReader) readSlice() ([]byte, error) {
	var fullLine []byte

	for {
		line, err := lr.r.ReadSlice('\n')
		fullLine = append(fullLine, line...)

		if err != nil {
			if err == bufio.ErrBufferFull {
				continue
			}

			if err == io.EOF {
				if len(fullLine) > 0 {
					return fullLine, nil
				}
				return nil, io.EOF
			}
			return nil, err
		}
		return fullLine, nil
	}
}
