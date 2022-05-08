/*
 * Created by Karthik Kolathumani
 * Created on Sat May 07 2022
 *
 */

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
