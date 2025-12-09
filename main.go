package main

import "fmt"

func main() {
	// s := []byte{65, 66, 67, 0, 3, 2, 68, 69, 70}
	encoded := []byte{65, 66, 0, 2, 2, 0, 2, 2}
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

	fmt.Println(string(decoded))
}
