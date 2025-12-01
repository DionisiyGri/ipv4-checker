package ipchecker

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExecute(t *testing.T) {
	tmp := filepath.Join(os.TempDir(), "ips.txt")
	content := `112.23.41.2
	1.2.3.4
	255.255.255.255
	1.2.3.4`

	err := os.WriteFile(tmp, []byte(content), 777)
	if err != nil {
		t.Fatalf("write tmp: %v", err)
	}
	defer os.Remove(tmp)

	res, err := Execute(tmp)
	if err != nil {
		t.Fatalf("counter error: %v", err)
	}

	if res.Lines != 4 {
		t.Fatalf("expected 4 lines, got %d", res.Lines)
	}
	if res.Unique != 3 {
		t.Fatalf("expected 3 unique, got %d", res.Unique)
	}
}
