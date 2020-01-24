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

# Packages in Go
- Each Go program begins execution in a package named main

## Key Pacakges

- strings : for manipulating strings values
- io : input and output functions
- bytes : for manipulating byte data. Provide similar functionality on byte data that strings does for string data
- os : supplies an interface that is platform independent for manipulating and managing files and folders
- path/filepath : similarly to os
- errors : for manipulating errors. It can allow you to create a custom error in addition to the built-in ones 
- container/list : Implements a doubly linked list
- hash : Hashes and cryptography
- encoding/gob : Allows you to enconde values. it allow for example create a network server that can handle client requests
- net/rpc/jsonrpc : For managing remote procedure calls. These packages provide a way to make methods available so they could be  called over a network


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


## Structs
```go
package main

import "fmt"

// You can think of a struct as a data structure
// made up of some fields of data
type  Rect struct { // struct name
	Width int
	Height int
}

func main() {

	r := Rect{7, 8} // creating a Rect with Width 7 and Height 8
	r.Width = 18 // we use period operator and the field name directly to set the value
	fmt.Println(r) // {18,8}

	// pointers and strcuts
	p := &r
	p.Width = 10
	fmt.Println(r) // {10,8}

	var (
		r1 = Rect {7,8} // type Rect
		r2 = Rect {Width: 4} // Height is implicitly 0
		r3 = Rect {} // Width and Height is implicitly 0
		p1 = Rect {7,8} // type * Rect
	)

	fmt.Println(r1,r2,r3,p1) 

}
 
```

## Maps
```go
package main

import "fmt"

// Maps are one of Go's advanced composite types
// They are similiar to arrays and slices being a collection of things
// But they are unordered
// Maps are a collection of key-value pairs

type Rect struct { // struct name
	Width  int
	Height int
}

func main() {
	// initializing a map with keys of type
	//string and values of type Rect
	var m = map[string]Rect{
		"Rect1": Rect{1, 2},
		"Rect2": Rect{4, 6},
	}
	fmt.Println(m)

	// we define a map 'w' using the make function
	w := make(map[string]int)

	w["Answer"] = 10
	fmt.Println("The value:", w["Answer"])

	w["Answer"] = 20
	fmt.Println("The value:", w["Answer"])

	// delete a value
	//  in this case the value is set to zero
	// because is int
	delete(w, "Answer")

	// for testing if a key is present
	// You assing two values on the left side of the expression
	v, ok := w["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}
```

## Advanced Functions
```go
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

```
## Methods
```go
package main

import (
	"fmt"
	"math"
)

type Root struct {
	A, B float64
}

/*
Methods are Functions declared with a Receiver.
Receiver is similar to a parameter. It can be a value or a pointer to a named or a struct type
Receiver is passes before the function name
*/

// this is a function that use a Pointer as a receiver
func (r *Root) Hyp() float64 {
	return math.Sqrt(r.A*r.A + r.B*r.B)
}

// we have used the type declaration to create a type Root1
type Root1 float64

// this is a function that use a value as a receiver
func (r Root1) Abs() float64 {
	if r < 0 {
		return float64(-r)
	}
	return float64(r)

}

func main() {

	r := Root{5, 6}

	// you can call de method Hyp() in your struct of type Root
	fmt.Println("C =", r.Hyp())

	// we create a new variable of type Root 1 and
	// it is initialized with the value of negative square root of 2
	r1 := Root1(-math.Sqrt2)

	// you can call the method Abs() in the Root1 type
	fmt.Println("C =", r1.Abs())
}
```

## Use Value or Pointer Receiver

### Value Receiver
- Receiver is map, func, or chan
- Receiver is an int (include all numeric types) or string
- Receiver is a small array or struct with no mutable fields or pointers

### Pointer Receiver
- Method needs to mutate receiver
- Receiver is a struct containing a synchronization field
- Receiver is large struct or array

## Interfaces

- Interfaces provides flexibility and abstraction
- It Specifies how values and pointers of specific type behave
- It specifies a method set, all methods for an interface type are the interface

```go
package main

import (
	"fmt"
	"math"
)

// defining a Interface with name Notifier
//  when the Interface has only one method is good practice
// to give interface a name containing an -er suffix
// Can specify numerous methods
type Notifier interface {
	Notify() error
}

// note
// the number of interface in Go's standard library with more than two methods are few and far between

//example,  we define a Calcer interface
type Calcer interface{ Calc() float64 }

type Square struct{ X, Y float64 }
type myFloat float64

// we define the Interface Calc to  work on two different types
// a  myFloat type 
func (f myFloat) Calc() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
// and  a Square type
func (s *Square) Calc() float64 {
	return math.Sqrt(s.X*s.X + s.Y*s.Y)
}

func main() {

	var c Calcer
	f := myFloat(-math.Sqrt2)
	s := Square{8, 6}
	c = &s

	// the Method employee in each case depends on the value type
	fmt.Println(c.Calc())
	c = f
	fmt.Println(c.Calc())

}
```

## Concurrency

- The idea behind concurrency is to be able to run a bunch of smaller tasks by simultaneously synchronizing access to shared memory. Tasks can run independenlty
- Go's implementation of concurrency is rooted back in the 1970s,back to communicating sequential processes or CSP.

## Concurrency Vs Parallelism

- Concurrency is a series of independently running processes
- Parallelism is a series of simultaneously running processes

```go
package main

import (
	"fmt"
)

func f(msg string) {

	for i := 0; i < 10; i++ {
		fmt.Println(msg, ":", i)
	}
}

func main() {
	// concurrency is implemented by using goroutines
	go f("value of i") // explicit goroutine

	// using a goroutine th main function doesn't have to wait
	// until f function returns

	var input string
	// the call to Scanln is made here because if it was not
	// the program would exit before the loop finished printing
	// Scanln waits for input essentially pausing the execution of func main()
	fmt.Scanln(&input)

	// the function main() itself is an implicit goroutine
}
```

## Channels

- Channels are used by goroutines for communication between them and to intelligently synchronize execution. For example, one goroutine can tell another via a channel that a particular task is complete.

- Channels are unbuffered (synchronous channel) or buffered. Channels are unbuffered by default

```go
package main

import (
	"fmt"
)

/*Unbuffered Channel will accept a "send" only if there is "receive"  waiting for the corresponding send value
By default sends and receives will be blocked until both sender and receiver are ready
*/
func sum(a []int, ch chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	// the sum is sent on the channel
	ch <- sum
}

func main() {
	a := []int{7, 0, -3, 5, 0, 4}

	// we create an integer channel
	ch := make(chan int)

	// we call goroutines passing the channel
	go sum(a[:len(a)/2], ch)
	go sum(a[len(a)/2:], ch)

	//the execution of the rechannel operations gets blocked
	// until goroutines write data to the channel
	x, y := <-ch, <-ch 

	fmt.Println(x, y, x+y)

	// buffered channels can accept a defined number of values even if
	// there isn't a corresponding receiver for those values
	c := make(chan int, 5) // channel of length of 5

	// we place the values 1 and 3 in the buffer
	c <- 1
	c <- 3

	// we receive them
	fmt.Println(<-c)
	fmt.Println(<-c)

	// you have to be careful not to overfill the buffer
	// if that occurs, you'll get a scary deadlock error
	// that states that all goroutines are asleep
}
```





