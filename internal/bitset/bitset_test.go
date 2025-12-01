package bitset_test

import (
	"testing"

	"github.com/DionisiyGri/ipv4-checker/internal/bitset"
)

func TestBitsetSetGet(t *testing.T) {
	b := bitset.New(10) // small bucket count for test
	if b.Get(5) {
		t.Fatalf("expected bit 5 to be 0 (clear)")
	}
	if !b.Set(5) {
		t.Fatalf("expected Set(5) true on first set")
	}
	if b.Set(5) {
		t.Fatalf("expected Set(5) false on second set")
	}
	if !b.Get(5) {
		t.Fatalf("expected bit 5 to be set")
	}
}
