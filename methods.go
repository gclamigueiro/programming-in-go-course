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

/*
Use Value or Pointer Receiver

Value Receiver
- Receiver is map, func, or chan
- Receiver is an int (include all numeric types) or string
- Receiver is a small array or struct with no mutable fields or pointers

Pointer Receiver
- Method needs to mutate receiver
- Receiver is a struct containing a synchronization field
- Receiver is large struct or array

*/
