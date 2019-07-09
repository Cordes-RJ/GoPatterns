package introChannels

import "fmt"

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
