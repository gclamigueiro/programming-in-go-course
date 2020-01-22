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
