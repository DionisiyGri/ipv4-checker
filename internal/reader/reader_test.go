package reader_test

import (
	"io"
	"os"
	"testing"

	"github.com/DionisiyGri/ipv4-checker/internal/reader"
)

func TestLineReader_Read(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantLines []string
		wantErr   error
	}{
		{
			name:      "single line with newline",
			input:     "192.168.0.1\n",
			wantLines: []string{"192.168.0.1\n"},
			wantErr:   io.EOF,
		},
		{
			name:      "single line without newline",
			input:     "10.0.0.1",
			wantLines: []string{"10.0.0.1"},
			wantErr:   io.EOF,
		},
		{
			name:      "multiple lines with newline",
			input:     "1.1.1.1\n2.2.2.2\n3.3.3.3\n",
			wantLines: []string{"1.1.1.1\n", "2.2.2.2\n", "3.3.3.3\n"},
			wantErr:   io.EOF,
		},
		{
			name:      "multiple lines last one without newline",
			input:     "8.8.8.8\n9.9.9.9",
			wantLines: []string{"8.8.8.8\n", "9.9.9.9"},
			wantErr:   io.EOF,
		},
		{
			name:      "empty file",
			input:     "",
			wantLines: []string{},
			wantErr:   io.EOF,
		},
		{
			name:      "single empty line",
			input:     "\n",
			wantLines: []string{"\n"},
			wantErr:   io.EOF,
		},
		{
			name:      "CRLF support",
			input:     "4.4.4.4\r\n5.5.5.5\r\n",
			wantLines: []string{"4.4.4.4\r\n", "5.5.5.5\r\n"},
			wantErr:   io.EOF,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmp, err := os.CreateTemp(t.TempDir(), "file.txt")
			if err != nil {
				t.Fatalf("temp file error: %v", err)
			}
			defer tmp.Close()

			if _, err := tmp.WriteString(tt.input); err != nil {
				t.Fatalf("write error %v", err)
			}
			if _, err := tmp.Seek(0, 0); err != nil {
				t.Fatalf("seek error %v", err)
			}

			lr, err := reader.New(tmp.Name(), 1<<32)
			if err != nil {
				t.Fatalf("failed to create reader")
			}

			var got []string
			for {
				line, err := lr.Read()
				if err == io.EOF {
					if len(line) > 0 {
						got = append(got, string(line))
					}
					break
				}
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				got = append(got, string(line))
			}

			if len(got) != len(tt.wantLines) {
				t.Fatalf("line count mismatch: got=%d want=%d lines: %q", len(got), len(tt.wantLines), got)
			}
			for i := range got {
				if got[i] != tt.wantLines[i] {
					t.Errorf("line[%d] mismatch: got=%q want=%q", i, got[i], tt.wantLines[i])
				}
			}
		})
	}
}
