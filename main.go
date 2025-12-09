package main

import (
	"fmt"
)

const lorem = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut non condimentum felis, ut viverra lorem. Cras laoreet felis non mi egestas, quis vestibulum turpis bibendum. Vivamus ex leo, cursus euismod leo eget, cursus condimentum velit. Sed aliquet, augue in gravida cursus, urna magna lobortis est, in lobortis orci ligula in lorem. Curabitur fringilla consequat mauris in elementum. Interdum et malesuada fames ac ante ipsum primis in faucibus. Suspendisse semper, ligula tempor gravida commodo, nulla mauris mollis metus, a fermentum dolor justo sed arcu. Nam vel varius erat. Sed vehicula aliquet magna. Sed facilisis sem sed ligula tristique mollis."

func main() {
	fmt.Println("Encoding...")
	fmt.Printf("Input bytes: %d\n", len([]byte(lorem)))
	encoded := encode([]byte(lorem))
	fmt.Printf("Encoded bytes: %d\n", len(encoded))
	fmt.Println("Decoding...")
	// s := []byte{65, 66, 67, 0, 3, 2, 68, 69, 70}
	// encoded := []byte{65, 66, 0, 2, 2, 0, 2, 2}
	decoded := decode(encoded)
	fmt.Println(string(decoded))
}

func encode(input []byte) []byte {
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

func decode(encoded []byte) []byte {
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
