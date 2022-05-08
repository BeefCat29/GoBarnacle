package main

import "fmt"

func main() {
	name := "Karthik Kolathumani"
	course := "Go Fundamentals"
	const dontChange = 186000 // the only way to declare constants
	fmt.Println("Hi", name, "your current course is", course)

	updateCoursePassByReference(&course)
	fmt.Println("Updated the course to ", course)

	course = updateCoursePassByValue(course)
	fmt.Println("Updated the course to ", course)
}

// Pass by value example
func updateCoursePassByValue(course string) string {
	course = "New Course Go advanced using pass by value" // Here we are not using := as we are assigning a mew value to an existing one
	return course
}

// Pass by reference
func updateCoursePassByReference(course *string) {
	*course = "New Course Go advanced using pass by reference"
}
