# Introduction
- Developed by Google in 2007, it was announced in 2009
- General purpose programming languaje
- Open source
- Clean, expressive, concise, and efficient syntax

## Key feautes
- Well-documented library
- Ideal for multicore applications and for network services 
- It is ideally suited for modular software design
- Compiles rapidly
- Supports run-time reflection includinf flexibility (this referes to the capability of a program to inspect or investigate its own structure) 

## Installing (windows)
  
  - download in https://golang.org/dl/ and install
  - check in the cmd using "go version" command
  
  
 # Hello World
- Create a file named hello.go

- Copy the code
```go
package main // package main is a special package, it tells that it's an executable program

import "fmt" // format package, contains functions for formatting output

// entry point for this program,
func main() { // important: curly brace has to be on the same line as your function declaration

	fmt.Println("Hello World") // print Hello World
}
```

- To compile and run it type `go run ./hello.go` in the cmd

# Variables

 ```go
 package main

import "fmt"

// variables could be declares in package scope
// using var variableName dataType
// or passing a value and the compiler infer the type
var hello = "Hello World"

// you can declare multiple variables at once.
var a1, a2 int

//Variables declared without a corresponding initialization are zero-valued.
//For example, the zero value for an int is 0. false for bool and "" for string

func main() {

	// you can use de := shorthand for declaring and initializing variables
	// but only in function scope
	bye := "Bye World"

	a1 = 1
	a2 = 2

	// declaring a constant value
	const notChange bool = true

	fmt.Println(a1, hello)
	fmt.Println(a2, bye)
}
```
# Functions

```go
 
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

```

# Parameters by 'value' and 'reference'

 ```go
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
```

# Operator 'new'

 ```go
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
```

# Flow Control Statements
 ```go
package main

import "fmt"
import "time"

func main() {

	// for statement
	for y := 1; y <= 5; y++ {
		// if/else statement
		if y%2 == 0 {
			fmt.Println(y, "even")
		} else {
			fmt.Println(y, "odd")
		}
	}
	
	//type switch 
	x := 4

	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	}

	// expression switch
	length := 9

	switch {
	case length < 4:
		fmt.Println("short")
	case length >= 4 && length < 8:
		fmt.Println("medium")
	case length >= 9:
		fmt.Println("large")
	}

	//switch with mutliple values
	switch time.Now().Weekday() {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("week day")
	default:
		fmt.Println("weekend")
	}
}
```

# Defer Statement
```go
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
```

# Arrays in go
```go
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
```

# Slices

## Declaring slices
```go
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
```

## Operations with slices
```go
package main

import "fmt"

func main() {

	slice := []int{7, 9, 10, 12}
	fmt.Println("slice =", slice)
	// using for loop with a slide
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%d] = %d\n", i, slice[i])

	}

	// example of [high:low] expresion
	fmt.Println("slice[1:4] =", slice[1:4]) //the high index is not inclusive
	fmt.Println("slice[:3] =", slice[:3]) // low index is omitted, so its value is zero
	fmt.Println("slice[1:4] =", slice[2:])  // high index is omitted, so Go assumes that the high index is the len of slice

	// example of range functions (it is like foreach)
	pow := []int{1,2,4,8,16,32,64}
	// range is a function that what it does the for loop to loop through each value within our slice
	for n,p := range pow {
		fmt.Printf("2**%d = %d\n", n, p)
	}

	// append function
	slice1 := []int{7,8,9}
	slice2 := append(slice1,10,11)
	fmt.Println(slice1, slice2)

	// copy function
	// the copy function only copy those items that will fit starting from index 0
	slice1 = []int{7,8,9}
	slice2 = make([]int,2)
	copy(slice2,slice1)
	fmt.Println(slice1, slice2)

	// if we do not initialize a slice, its len and capacity will be zero
	// an if we compare with nil, the comparation returns true
	var slice3 []int
	fmt.Println(slice3,len(slice3),cap(slice3))
	if slice3 == nil {
	fmt.Println("nil")
	}		

}
 
```
