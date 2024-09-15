package variablelengthquantity

import (
	"fmt"
)

func EncodeVarint(input []uint32) []byte {
	result := make([]byte, 0)
	for _, v := range input {
		encoded := encodeInt(v)
		result = append(result, encoded...)
	}
	return result
}

func encodeInt(n uint32) []byte {
	if n == 0 {
		return []byte{0}
	}

	var result []byte
	for n > 0 {
		// Get 7 rightmost bits
		part := n & 0x7f
		// Shift 7 bits to the right
		n >>= 7
		result = append(result, byte(part))
	}

	// Reverse
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	// All bytes except the last one should have MSB set to 1
	for i := 0; i < len(result)-1; i++ {
		result[i] |= 0x80
	}
	return result
}

func DecodeVarint(input []byte) ([]uint32, error) {
	var (
		result       []uint32
		current      uint32
		fullSequence bool
	)

	for _, octet := range input {
		// Add the 7 rightmost bits to the temporary accumulator
		current = (current << 7) | uint32(octet&0x7f)

		// Not last octet? Move on
		if octet&0x80 != 0 {
			continue
		}

		// Last octet handling
		result = append(result, current)
		current = 0
		fullSequence = true
	}

	if !fullSequence {
		return nil, fmt.Errorf("incomplete sequences")
	}

	return result, nil
}
