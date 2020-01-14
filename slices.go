package main

import "fmt"

// slices are segment of an array (always are associated with an underlying array )
// they are indexable and length can vary
func main() {

	// creating a slice of int
	mslice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// takin a slice of antoher slice using [high:low] expresion
	s := mslice[0:2]

	fmt.Println(mslice,s)

	// creating a slice with 'make' method
	// parameters     type, len, capacity
	otherSlice := make([]string, 10, 20)
	otherSlice[0] = "Testing"

	fmt.Println(otherSlice)

}
