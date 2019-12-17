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
 
  

 
