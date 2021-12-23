package variablelengthquantity

func EncodeVarint(input []uint32) (encoded []byte) {
	var remainder uint32

	for _, val := range input {
		var base128Val []byte

		if val == 0 {
			encoded = append(encoded, byte(0))
		}

		// Convert val to base128 representation
		for val > 0 {
			val, remainder = val/128, val%128
			// Prepend remainder to base128Val
			base128Val = append([]byte{byte(remainder)}, base128Val...)
		}

		for i, val := range base128Val {
			// Add a leading bit if this is not the last byte
			if i < len(base128Val)-1 {
				val = 0b10000000 | val
			}
			encoded = append(encoded, val)
		}
	}
	return encoded
}

func DecodeVarint(input []byte) ([]uint32, error) {
	panic("Please implement the EncodeVarint function")
}
