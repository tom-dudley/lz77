package main

import (
	"fmt"

	"github.com/tom-dudley/lz77"
)

const lorem = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut non condimentum felis, ut viverra lorem. Cras laoreet felis non mi egestas, quis vestibulum turpis bibendum. Vivamus ex leo, cursus euismod leo eget, cursus condimentum velit. Sed aliquet, augue in gravida cursus, urna magna lobortis est, in lobortis orci ligula in lorem. Curabitur fringilla consequat mauris in elementum. Interdum et malesuada fames ac ante ipsum primis in faucibus. Suspendisse semper, ligula tempor gravida commodo, nulla mauris mollis metus, a fermentum dolor justo sed arcu. Nam vel varius erat. Sed vehicula aliquet magna. Sed facilisis sem sed ligula tristique mollis."

func main() {
	fmt.Println("Encoding...")
	fmt.Printf("Input bytes: %d\n", len([]byte(lorem)))
	encoded := lz77.Encode([]byte(lorem))
	fmt.Printf("Encoded bytes: %d\n", len(encoded))
	fmt.Println("Decoding...")
	decoded := lz77.Decode(encoded)
	fmt.Println(string(decoded))
}
