package main

import "fmt"

// You can think of a struct as a data structure
// made up of some fields of data
type Rect struct { // struct name
	Width  int
	Height int
}

func main() {

	r := Rect{7, 8} // creating a Rect with Width 7 and Height 8
	r.Width = 18    // we use period operator and the field name directly to set the value
	fmt.Println(r)  // {18,8}

	// pointers and strcuts
	p := &r
	p.Width = 10
	fmt.Println(r) // {10,8}

	var (
		r1 = Rect{7, 8}     // type Rect
		r2 = Rect{Width: 4} // Height is implicitly 0
		r3 = Rect{}         // Width and Height is implicitly 0
		p1 = Rect{7, 8}     // type * Rect
	)

	fmt.Println(r1, r2, r3, p1)

}
