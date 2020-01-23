package main

import (
	"fmt"
	"time"
)

func msg() {
	// simulate a time-consuming process (5 seconds)
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Millisecond * 1000)
		if i > 3 {
			fmt.Println(i, "seconds... yawn")
		} else {
			fmt.Println(i, "seconds")
		}
	}
}

func main() {
	// call  msg function but main doesn't wait
	// for it to continue
	go msg()

	// let's print to the console when we're finished
	fmt.Println("Finished")

	// it is added a little bit more delay in here for
	// waiting for func msg() to finish
	time.Sleep(time.Millisecond * 7000)

}
