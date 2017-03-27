// Package plessey provides "UK" Plessey support
package plessey

import (
	"strconv"
	"strings"
)

// Bits .
type Bits []uint8

// Plessey table
var table = Bits{
	0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0,
	0, 0, 1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 0,
	0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 1, 1, 1, 0, 1,
	0, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1,
}

// Plessey CRC
var crc = Bits{1, 1, 1, 1, 0, 1, 0, 0, 1}

// Checksum .
func (bits Bits) Checksum() Bits {
	var checkptr = make(Bits, len(bits))
	copy(checkptr, bits)

	for i := 0; i < len(checkptr)-8; i++ {
		if checkptr[i] == 1 {
			for j := 0; j < len(crc); j++ {
				checkptr[i+j] ^= crc[j]
			}
		}
	}
	return checkptr[len(checkptr)-8:]
}

// ToBits .
func ToBits(barcode string) (Bits, error) {
	var ret = Bits{}
	barcode = strings.ToUpper(barcode)

	for _, c := range barcode {
		if c == 'X' {
			c = 'A'
		}
		idx, err := strconv.ParseInt(string(c), 16, 8)
		if err != nil {
			return ret, err
		}

		ret = append(ret, table[idx*4:idx*4+4]...)
	}

	return ret, nil
}

// Bits to string
func (bits Bits) String() string {
	const hex = "0123456789ABCDEF"
	ret := ""
	for i := 0; i < len(bits)/4; i++ {
		slice := bits[i*4 : i*4+4]
		for j := 0; j < len(table)/4; j++ {
			if string(table[j*4:j*4+4]) == string(slice) {
				ret += string(hex[j])
				continue
			}
		}
	}
	return ret
}
