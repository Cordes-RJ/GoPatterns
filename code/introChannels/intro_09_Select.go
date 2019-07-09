package introChannels

import (
	"fmt"
	"time"
)

// SelectSimple09 is a simple example of how to use select statements
func SelectSimple09() {
	c0 := make(chan int)
	c1 := make(chan int)
	var c0ct, c1ct int
	go func() {
		for i := 0; i < 20; i++ {
			select {
			case <-c0:
				c0ct++
			case <-c1:
				c1ct++
			}
		}
	}()
	for i := 0; i < 20; i++ {
		x := i
		if i%2 == 0 {
			c0 <- x
		} else {
			c1 <- x
		}
	}
	close(c0)
	close(c1)
	fmt.Printf("[%d]vs[%d]", c0ct, c1ct)
}

func SelectTimeOut09() {
	c0 := make(chan int, 10)
	c1 := make(chan int, 10)
	var c0ct, c1ct int
	go func() {
		timestart := time.Now()
		for i := 0; i < 1000; i++ {
			select {
			case <-c0:
				c0ct++
			case <-c1:
				c1ct++
			case <-time.After(2 * time.Millisecond):
				fmt.Printf("Timed-out at i[%d]", i)
				goto End
			}
		}
	End:
		fmt.Printf("%v", time.Since(timestart))
	}()
	for i := 0; i < 1000; i++ {
		x := i
		if i%2 == 0 {
			select {
			case c0 <- x:
				time.Sleep(1 * time.Nanosecond)
			default:
				break
			}
		} else {
			select {
			case c1 <- x:
				time.Sleep(1 * time.Nanosecond)

			default:
				break
			}
		}
	}
	close(c0)
	close(c1)
	fmt.Printf("...[%d]vs[%d]", c0ct, c1ct)
}
