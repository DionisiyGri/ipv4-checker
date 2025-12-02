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
	// Read lines until EOF
	for {
		line, err := lr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return Result{}, fmt.Errorf("reading line: %w", err)
		}
		res.Lines++

		ipNum, err := ipToUint32(trim(line))
		if err != nil {
			log.Printf("cant convert ip [%s] to uint: %v", line, err)
			continue
		}
		if bs.Set(ipNum) {
			res.Unique++
		}
	}
	return res, nil
}

// Trim spaces and newline characters
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
	var octets [4]uint32
	idx := 0

	for _, c := range b {
		if c == '.' {
			idx++
			if idx > 3 {
				return 0, fmt.Errorf("invalid ip")
			}
			continue
		}
		if c < '0' || c > '9' {
			return 0, fmt.Errorf("invalid char")
		}

		octets[idx] = octets[idx]*10 + uint32(c-'0')
		if octets[idx] > 255 {
			return 0, fmt.Errorf("invalid octet")
		}
	}

	if idx != 3 {
		return 0, fmt.Errorf("invalid ip")
	}

	return octets[0]<<24 | octets[1]<<16 | octets[2]<<8 | octets[3], nil
}
