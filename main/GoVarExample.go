package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// These are variables declared at the package level
// These may or may not be used since they are the package level
var (
	name, course        string
	module              int
	clip                int
	firstName, lastName = "Karthik", "Kolathumani"
	age                 = 30
	convertToInt        = "99"
)

func main() {

	// Variable assignments and usage
	fmt.Println("Name and course are set to", name, "and", course, ".")
	fmt.Println("Module and clip are set to", module, "and", clip, ".")
	fmt.Println("Name is of type", reflect.TypeOf(name))
	fmt.Println("Module is of type", reflect.TypeOf(module))
	fmt.Println("Module is string type -", reflect.TypeOf(reflect.TypeOf(module))) //*reflect.rtype
	fmt.Println("FirstName and LastName", firstName, lastName)

	//Type conversions Atoi
	iModule, err := strconv.Atoi(convertToInt)
	if err == nil {
		clip := 1 + iModule // This assignment only works within functions
		fmt.Println("Atoi should be 100 and is", clip)
	}

	courseComplete := false // This varible needs to be used since its declared inside a function body
	fmt.Println("Was the course completed?", courseComplete)

	// Getting aquainted with pointers
	fmt.Println("Address of the variable courseComplete of type", reflect.TypeOf(courseComplete), "is", &courseComplete)

	var ptr *bool = &courseComplete
	fmt.Println("Address of the variable courseComplete using pointer is", ptr, "and value is", *ptr)
}
