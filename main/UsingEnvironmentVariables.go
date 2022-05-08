package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("The current logged in user is", os.Getenv("USER"))

	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}
