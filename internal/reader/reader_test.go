package reader

import (
	"os"
	"testing"
)

func TestReader(t *testing.T) {
	content := `a
	b
	c
	f`

	tmp, err := os.CreateTemp(t.TempDir(), "file.txt")
	if err != nil {
		t.Fatalf("temp file error: %v", err)
	}
	defer tmp.Close()

	if _, err := tmp.WriteString(content); err != nil {
		t.Fatalf("write error %v", err)
	}
	if _, err := tmp.Seek(0, 0); err != nil {
		t.Fatalf("seek error %v", err)
	}

	lr, err := New(tmp.Name(), 1<<32)
	if err != nil {
		t.Fatalf("create new reader error: %v", err)
	}

	lines, errCh := lr.Read()

	var parsed []string
	for line := range lines {
		parsed = append(parsed, line)
	}

	var readErr error
	select {
	case readErr = <-errCh:
	default:
	}

	if readErr != nil {
		t.Fatalf("reader returned error: %v", readErr)
	}

	expected := 4
	if len(parsed) != expected {
		t.Fatalf("expected %d lines, got %d", expected, len(parsed))
	}
}
