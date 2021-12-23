package variablelengthquantity

import "errors"

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

func DecodeVarint(input []byte) (decoded []uint32, e error) {
	var decodedInt uint32

	if input[len(input)-1]&128 != 0 {
		return nil, errors.New("incomplete sequence")
	}

	for _, val := range input {
		decodedInt = 128*decodedInt + uint32(127&val)
		if 128&val == 0 {
			decoded = append(decoded, decodedInt)
			decodedInt = 0
		}
	}
	return decoded, nil
}
