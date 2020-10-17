// Package variablelengthquantity is solution for problem Variable Length Quantity.
package variablelengthquantity

import (
	"fmt"
)

// EncodeVarint encodes series of uint32 using VLQ-encoding.
func EncodeVarint(input []uint32) []byte {
	var encoded []byte
	for _, v := range input {
		encoded = append(encoded, encode(v)...)
	}
	return encoded
}

func encode(n uint32) []byte {
	if n == 0 {
		return []byte{0}
	}

	var result []byte
	for n != 0 {
		// prepend last 8 bits to result
		result = append([]byte{byte(n & 0x7f)}, result...)
		// shift 7 times to remove last 7 bits appended to result earlier
		n >>= 7
		// if this is not the last byte, set the MSB to 1
		if len(result) > 1 {
			result[0] |= 0x80
		}
	}
	return result
}

// DecodeVarint decodes VLQ-encoded bytes to uint32.
func DecodeVarint(input []byte) ([]uint32, error) {
	var results []uint32
	var result uint32
	var complete bool
	for _, b := range input {
		if b >= 0x80 {
			complete = false	
		}

		result = result << 7 | uint32(b & 0x7f)
		if b < 0x80 {
			results = append(results, result)
			result = 0
			complete = true
		}
	}

	if !complete {
		return nil, fmt.Errorf("invalid sequence")
	}
	return results, nil
}
