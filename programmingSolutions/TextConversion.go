/*
 * Created by Karthik Kolathumani
 * Created on Mon May 09 2022
 *
 */

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	// Title case example
	author := "karthik kolathumani"
	course := "go programming"

	fmt.Println(converterString(author, course))
	fmt.Println(convertStringInt(author, strconv.Itoa(math.MaxInt)))
}

// Use of a function returning two values
func converterString(author, course string) (authorCase, courseCase string) {
	authorTitleCase := strings.Title(author)
	courseUpperCase := strings.ToUpper(course)

	return authorTitleCase, courseUpperCase
}

// Use of a function returing multiple values with error
func convertStringInt(author, authorId string) (authorCase string, authorIdConv int, err error) {
	authorTitleCase := strings.Title(author)
	authorIdInt, err := strconv.Atoi(authorId)
	if err == nil {
		return authorTitleCase, authorIdInt, err
	} else {
		return authorTitleCase, math.MinInt, err
	}
}
