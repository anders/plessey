// Package plessey provides "UK" Plessey support
package plessey

// Bits .
type Bits []uint8

var table = [16]Bits{
	Bits{0, 0, 0, 0}, Bits{1, 0, 0, 0}, Bits{0, 1, 0, 0}, Bits{1, 1, 0, 0},
	Bits{0, 0, 1, 0}, Bits{1, 0, 1, 0}, Bits{0, 1, 1, 0}, Bits{1, 1, 1, 0},
	Bits{0, 0, 0, 1}, Bits{1, 0, 0, 1}, Bits{0, 1, 0, 1}, Bits{1, 1, 0, 1},
	Bits{0, 0, 1, 1}, Bits{1, 0, 1, 1}, Bits{0, 1, 1, 1}, Bits{1, 1, 1, 1},
}

var fromASCII = map[rune]uint8{
	'0': 0, '1': 1, '2': 2, '3': 3, '4': 4,
	'5': 5, '6': 6, '7': 7, '8': 8, '9': 9,

	'a': 0xa, 'b': 0xb, 'c': 0xc, 'd': 0xd, 'e': 0xe, 'f': 0xf,
	'A': 0xa, 'B': 0xb, 'C': 0xc, 'D': 0xd, 'E': 0xe, 'F': 0xf,
}

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

	for i := 0; i < len(barcode); i++ {
		c := rune(barcode[i])
		idx, ok := fromASCII[c]
		if !ok {
			continue
		}
		ret = append(ret, table[idx]...)
	}

	return ret
}

// Bits to string
func (bits Bits) String() string {
	const hex = "0123456789ABCDEF"
	ret := ""
	for i := 0; i < len(bits)/4; i++ {
		slice := bits[i*4 : i*4+4]
		for j := 0; j < len(table); j++ {
			if string(table[j]) == string(slice) {
				ret += string(hex[j])
				continue
			}
		}
	}
	return ret
}
