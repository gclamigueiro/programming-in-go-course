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
