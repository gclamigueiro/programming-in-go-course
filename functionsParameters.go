package main

import "fmt"

// in this function the value of 'y' is passed as a copy (by value)
// it means that if you change 'y' inside the function
// the value is set in the local function only
func fun(y int) {
	y = 0
}

// in this case the parameter is a pointer (by reference)
// if we change the value that the pointer points
// the change will be reflected in the main function
func funWithPtr(yPtr *int) {

	// the * basically means that we're providing acces
	// to the underlying value that the pointer points to
	// in this case we are store the value 0 in the memory
	// location pointer to by "y" pointer
	*yPtr = 0
}

func main() {

	// declaring a variable
	y := 10

	// passing variable as a copy
	fun(y)

	// in this case, it will print 10
	fmt.Println(y)

	// using the & operator, what this does is it basically causes
	// us to send the memory location (a pointer)
	funWithPtr(&y)

	// in this case, it will print 0
	fmt.Println(y)
}
