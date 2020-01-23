package main

import (
	"fmt"
)

/*unbuffered Channel will accept a "send" only if there is "receive"  waiting for the corresponding send value
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
