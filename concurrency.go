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
