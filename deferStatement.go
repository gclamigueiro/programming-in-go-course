package main

import "fmt"

// a defer statement defers de execution of a function until the function
// that is contained under returns
func main() {

	defer fmt.Println("World")
	fmt.Println("Hello") // Occurs first

	// Deferred functions calls pushed onto a list
	// Deferred calls executed last-in-first-out (stack)
	fmt.Println("counting")

	for x := 0; x < 5; x++ {
		defer fmt.Println(x)
	}

	fmt.Println("finished")

	// deferred functions run even if a runtime panics occurs
	defer func() {
		str := recover() // recover function is a built-in to halt the panic.
		// it returns the value that was passed to panic
		fmt.Println(str)
	}()

	panic("PANIC") // panic is a function, wich is a way to create a runtime error

}
