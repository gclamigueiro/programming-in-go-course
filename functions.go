package main

import "fmt"

/* function declaration*/
/* important: the opening curly brace has to be
on the same line as the function declaration */
func add(x int, y int) int {
	return x + y
}

/*you can declare the parameters in this way*/
func sub(x, y int) int {
	return x - y
}

/*you can have multiple result returned */
func mult(a, b string) (string, string) {
	return a, b
}

func main() {

	a1 := 3
	a2 := 2
	fmt.Println(add(a1, a2))

	x, y := mult("Hello", "World")
	fmt.Println(x, y)

}
