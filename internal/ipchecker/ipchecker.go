package ipchecker

import (
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/DionisiyGri/ipv4-checker/internal/bitset"
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

	//allocate bitset 2^32 bits / 64 = 2^26 uints (~512Mb)
	const totalBits = uint64(1) << 32
	const bitsPerBucket = uint64(64)
	buckets := totalBits / bitsPerBucket
	bs := bitset.New(buckets)

	var res Result
	for {
		lines, err := lr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return Result{}, fmt.Errorf("reading line: %w", err)
		}
		res.Lines++

		ipNum, err := ipToUint32(trim(lines))
		if err != nil {
			log.Printf("cant convert ip [%s] to uint: %w", lines, err)
			continue
		}
		if bs.Set(ipNum) {
			res.Unique++
		}
	}
	return res, nil
}

// trim trims lines, spaces with no allocations
func trim(b []byte) []byte {
	start := 0
	end := len(b)

	// Trim left
	for start < end {
		c := b[start]
		if c == ' ' || c == '\t' || c == '\r' || c == '\n' {
			start++
			continue
		}
		break
	}

	// Trim right
	for start < end {
		c := b[end-1]
		if c == ' ' || c == '\t' || c == '\r' || c == '\n' {
			end--
			continue
		}
		break
	}

	return b[start:end]
}

// ipToUint32 converts ip into uint32
func ipToUint32(b []byte) (uint32, error) {
	var p [4]uint32
	part := 0

	for _, c := range b {
		if c == '.' {
			part++
			if part > 3 {
				return 0, fmt.Errorf("invalid ip")
			}
			continue
		}
		if c < '0' || c > '9' {
			return 0, fmt.Errorf("invalid char")
		}

		p[part] = p[part]*10 + uint32(c-'0')
		if p[part] > 255 {
			return 0, fmt.Errorf("invalid octet")
		}
	}

	if part != 3 {
		return 0, fmt.Errorf("invalid ip")
	}

	return p[0]<<24 | p[1]<<16 | p[2]<<8 | p[3], nil
}
