package ipchecker_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DionisiyGri/ipv4-checker/internal/ipchecker"
)

func TestExecute(t *testing.T) {
	tests := []struct {
		name       string
		content    string
		wantLines  uint64
		wantUnique uint64
		wantErr    bool
	}{
		{
			name:       "valid unique IPs",
			content:    "1.1.1.1\n8.8.8.8\n9.9.9.9\n",
			wantLines:  3,
			wantUnique: 3,
		},
		{
			name:       "duplicates counted once",
			content:    "4.4.4.4\n4.4.4.4\n4.4.4.4\n",
			wantLines:  3,
			wantUnique: 1,
		},
		{
			name:       "mix valid/invalid/empty",
			content:    "145.67.23.4\ninvalid\n\n8.8.8.8\njunk\n",
			wantLines:  5,
			wantUnique: 2, // two valid
		},
		{
			name:       "empty file",
			content:    "",
			wantLines:  0,
			wantUnique: 0,
		},
		{
			name:       "ipv6 ignored",
			content:    "2001:db8::1\n8.8.8.8\n",
			wantLines:  2,
			wantUnique: 1, // only ipv4 counted
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tmp := filepath.Join(t.TempDir(), "test.txt")
			if err := os.WriteFile(tmp, []byte(tt.content), 0o644); err != nil {
				t.Fatalf("failed writing temp file: %v", err)
			}

			res, err := ipchecker.Execute(tmp)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Execute() error = %v wantErr %v", err, tt.wantErr)
			}

			if res.Lines != tt.wantLines {
				t.Errorf("Lines = %d want %d", res.Lines, tt.wantLines)
			}

			if res.Unique != tt.wantUnique {
				t.Errorf("Unique = %d want %d", res.Unique, tt.wantUnique)
			}
		})
	}
}

func TestExecute_EmptyPath(t *testing.T) {
	_, err := ipchecker.Execute("")
	if err == nil {
		t.Fatal("expected error for empty path")
	}
}
