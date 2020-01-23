package main

import (
	"fmt"
	"math"
)

/*
- Interfaces provides flexibility and abstraction
- It Specifies how values and pointers of specific type behave
- It specifies a method set, all methods for an interface type are the interface
*/

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
