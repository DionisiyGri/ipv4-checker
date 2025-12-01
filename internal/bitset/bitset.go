package bitset

// Bitset is a wrapper around []uint64 to store 2^32bits
// dont use zero value, call New to allocate
type Bitset struct {
	b []uint64
}

// New allocates buckets uint64s
func New(buckets uint64) *Bitset {
	return &Bitset{b: make([]uint64, buckets)}
}

// Set returns true if the bit was previously 0 (eg. newly set)
func (bs *Bitset) Set(n uint32) bool {
	idx := uint64(n) / 64    // which bucket
	pos := uint64(n % 64)    // position in bucket (0...63)
	mask := uint64(1) << pos // mask with single bit at pos
	cur := bs.b[idx]         // current bucket value

	if cur&mask != 0 {
		return false
	}
	bs.b[idx] = cur | mask
	return true
}

// Get returns true bit is set
func (bs *Bitset) Get(n uint32) bool {
	idx := uint64(n) / 64
	pos := uint64(n % 64)
	mask := uint64(1) << pos
	return (bs.b[idx] & mask) != 0
}
