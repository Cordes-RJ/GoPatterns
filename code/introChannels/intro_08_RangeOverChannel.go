package introChannels

import "fmt"

// RangeOverChan08 shows how one might use a channel's closure as a means of
// signaling between routines
func RangeOverChan08() {
	intPipe := make(chan int) // 1
	go func() {
		defer close(intPipe)
		for i := 1; i <= 5; i++ {
			intPipe <- (i * 10) // 2
		}
	}()
	for v := range intPipe { // 3
		fmt.Printf("[%d]", v)
	}
}

// [1] Make Channel
// [2] push to pipe until index passes 5
// [3] This range function will wait for values to pass through intpipe until
// intpipe is closed, effectively creating a wait() function without use of a
// WaitGroup
