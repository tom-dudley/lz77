package lz77

// Encode with crude LZ77
func Encode(input []byte) []byte {
	encoded := []byte{}
	for i := 0; i < len(input); i++ {
		lengths := []int{}
		for j := 0; j < i; j++ {
			length := 0
			for {
				if j+length == len(input) || i+length == len(input) {
					lengths = append(lengths, length)
					break
				}
				if input[j+length] == input[i+length] {
					length++
				} else {
					lengths = append(lengths, length)
					break
				}
			}
		}
		maxLength := 0
		maxLengthIndex := -1
		for j := i - 1; j >= 0; j-- {
			distance := i - j
			length := lengths[j]
			// We only allow lengths up to 255 (1 byte per length)
			if distance > 255 {
				break
			}
			if length > maxLength && length < 255 {
				maxLength = length
				maxLengthIndex = j
			}
		}

		if maxLength > 3 {
			encoded = append(encoded, 0x0)
			encoded = append(encoded, byte(i-maxLengthIndex))
			encoded = append(encoded, byte(maxLength))
			i += maxLength - 1

		} else {
			encoded = append(encoded, input[i])
		}
	}

	return encoded
}

// Decode with crude LZ77
func Decode(encoded []byte) []byte {
	decoded := []byte{}
	decodedIndex := 0

	for i := 0; i < len(encoded); i++ {
		if encoded[i] == 0 {
			lookback := encoded[i+1]
			length := encoded[i+2]
			for j := 0; j < int(length); j++ {
				char := decoded[decodedIndex-int(lookback)]
				decoded = append(decoded, char)
				decodedIndex++
			}
			i += 2
		} else {
			decoded = append(decoded, encoded[i])
			decodedIndex++
		}
	}
	return decoded
}
