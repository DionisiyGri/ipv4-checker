package ipchecker

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/DionisiyGri/ipv4-checker/internal/reader"
)

type Result struct {
	Lines  uint64
	Unique uint64
}

func Execute(path string) (Result, error) {
	if path == "" {
		return Result{}, errors.New("empty path")
	}

	lr, err := reader.New(path, 1<<20) // 1mb buffer
	if err != nil {
		return Result{}, fmt.Errorf("open file: %w", err)
	}
	defer lr.Close()

	//allocate bitset 2^32 bits / 64 = 2^26 uints
	const totalBits = uint64(1) << 32
	const bitsPerBucket = uint64(64)
	buckets := totalBits / bitsPerBucket

	// check if it is possible to allocate needed amount of memory
	var bitset []uint64
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Fprintf(os.Stderr, "failed to allocate %d uint64 buckets, ~%d bytes", buckets, buckets*8)
				os.Exit(1)
			}
		}()
		bitset = make([]uint64, buckets)
	}()

	var res Result
	linesCh, errCh := lr.Read()
	for ln := range linesCh {
		res.Lines++
		if ln == "" {
			continue
		}

		ip := net.ParseIP(ln)
		if ip == nil {
			continue
		}

		ip4 := ip.To4()
		if ip4 == nil {
			continue
		}

		//convert 4 bytes ipv4 into int32 value (a.b.c.d -> 0xAABBDDC)
		ipNum := binary.BigEndian.Uint32(ip4)
		if setBit(bitset, ipNum) {
			res.Unique++
		}
	}

	select {
	case err := <-errCh:
		if err != nil {
			return res, fmt.Errorf("read error: %w", err)
		}
	default:
	}

	return res, nil
}

func setBit(bitset []uint64, n uint32) bool {
	idx := uint64(n) / 64    // which bucket
	pos := uint64(n % 64)    // position in bucket (0...63)
	mask := uint64(1) << pos // mask with single bit at pos
	cur := bitset[idx]       // current bucket value

	if cur&mask != 0 {
		return false
	}
	bitset[idx] = cur | mask
	return true
}
