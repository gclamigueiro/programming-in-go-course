package main

import "fmt"

// variadic function
// the ... tells the compiler that there could be
// zero or more parameters, but we aren't sure how many
func add(args ...int) int {
	total := 0

	// range allows the for loop to loop through the entire
	// range of arguments
	for i, n := range args {
		print(i)
		total += n
	}

	return total
}

// a function that returns a function
func fibonacci() func() int {
	x := 0
	y := 1
	fmt.Println(x)
	// it returs a closure function that use the values of x and y
	return func() int {
		x, y = y, x+y
		return x
	}

}

func main() {
	// calling variadic function
	fmt.Println(add(8, 13, 4, 5, 3))

	// closure functions
	// a closure function can access local variables
	a := 4
	decrement := func() int {
		a--
		return a
	}

	fmt.Println(decrement())
	fmt.Println(decrement())

	// calling fibonacci function
	fmt.Println("Fibonacci")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())

	}

}
