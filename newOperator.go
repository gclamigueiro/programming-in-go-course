package main

import "fmt"

func funWithPtr(yPtr *int) {
	// changing the value
	*yPtr = 5
}

func main() {

	// we use the operator 'new' to obtain a pointer
	// you pass a type as an argument and then it allocates
	// enough  memory for a value of the type passed to it and
	// returns a pointer to that memory location
	y := new(int)

	// calling the function sending a memory address
	funWithPtr(y)

	// in this case, it will print 5
	fmt.Println(*y)
}
