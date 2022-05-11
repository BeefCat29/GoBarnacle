package main

import "fmt"

func main() {
	size := 10
	array := make([]int, size, size)
	// Init
	for i := 0; i < 10; i++ {
		array[i] = i
	}

	slicedArray := array[2:10]

	fmt.Println(len(array), cap(array))

	fmt.Println("Printing the sliced array")
	for _, i := range slicedArray {
		fmt.Println(i)
	}

	//try mutate the sliced array??
	slicedArray[2] = 10

	fmt.Println("Original array -", array[2], "Sliced array -", slicedArray[2])

}
