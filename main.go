package main

import "fmt"

func main() {
	s := []byte{65, 66, 67, 0, 3, 2, 68, 69, 70}
	for i := 0; i < len(s); i++ {
		if s[i] == 0 {
			lookback := s[i+1]
			length := s[i+2]
			for j := 0; j < int(length); j++ {
				char := s[i-int(lookback)+j]
				fmt.Println(string(char))
			}
			i += 2
		} else {
			fmt.Println(string(s[i]))
		}
	}
}
