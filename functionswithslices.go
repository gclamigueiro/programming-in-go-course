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
 