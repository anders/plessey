// Package plessey provides "UK" Plessey support
package plessey

import "strings"

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
	var temp = make(Bits, len(bits)+8)
	copy(temp, bits)

	for i := 0; i < len(bits); i++ {
		if temp[i] == 1 {
			for j := 0; j < len(crc); j++ {
				temp[i+j] ^= crc[j]
			}
		}
	}

	return temp[len(temp)-8:]
}

// ToBits .
func ToBits(barcode string) Bits {
	var ret = Bits{}
	barcode = strings.ToUpper(barcode)

	for _, c := range barcode {
		// idx: from ASCII 0-9, A-Z to 0-16
		idx := 0

		if c == 'X' {
			c = 'A'
		}

		if c >= '0' && c <= '9' {
			idx = int(c - '0')
		} else if c >= 'A' && c <= 'Z' {
			idx = int(c - 'A')
		} else {
			continue
		}

		// skip anything above "F""
		if !(c >= '0' && c <= '9') && c > 'F' {
			continue
		}

		ret = append(ret, table[idx*4:idx*4+4]...)
	}

	return ret
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
