package main

import (
	"fmt"
)

func decode(encoded []int, first int) []int {
	decoded := make([]int, len(encoded)+1)
	decoded[0] = first
	for i := 1; i < len(decoded); i++ {
		decoded[i] = encoded[i-1] ^ decoded[i-1]
	}
	return decoded
}

func main() {
	encoded := [4]int{6, 2, 7, 3}
	first := 4
	decoded := decode(encoded[:], first)

	for i := 0; i < len(decoded); i++ {
		fmt.Println(decoded[i])
	}
}
