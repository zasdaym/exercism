package variablelengthquantity

func EncodeVarint(input []uint32) []byte {
	var result []byte
	for i := 0; i < len(input); i++ {
		encoded := encode(input[i])
		result = append(result, encoded...)
	}
	return result
}

// encode n with variable length encoding.
// Only works for number between 0 and 127.
func encode(n uint32) []byte {
	var result []byte
	for n > 0 {
		// get the last 7 right bits
		encoded := n & 0x7f
		// shift n to the right by 7 bits
		n >>= 7
		// if there are more bits to encode, set the 8th bit to 1
		if n > 0 {
			encoded |= 0x80
		}
		result = append(result, byte(encoded))
	}
	return result
}

func DecodeVarint(input []byte) ([]uint32, error) {
	panic("Please implement the EncodeVarint function")
}
