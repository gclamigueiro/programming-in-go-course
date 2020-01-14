package main

import "fmt"

// in Go is more common the use of slices, which are based on arrays
// this is a small example of array
func main() {

	//declaring an array of string
	var z [3]string
	z[0] = "Hello"
	z[1] = " "
	z[2] = "World"

	fmt.Println(z) // Hello World

	// declarin an array of int
	var y [3]int
	y[2] = 4

	fmt.Println(y) // 0 0 4

}
