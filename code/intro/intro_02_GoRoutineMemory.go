package intro

import (
	"fmt"
	"sync"
)

// If we look at the function below:
func intro02memoryA() {
	var wg sync.WaitGroup // this is a WaitGroup
	for i, item := range []string{"a", "b", "c", "d"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("| $%s : %d |", item, i)
		}()
	}
	wg.Wait()
}

// We'd expect an output of a:1, b:2 ...
// but instead, I get d:3 d:3 d:3...
// You may also get an error on the line, suggesting
// a loop variable is captured by the func literal-
// because the Goroutines can run at ANY point in the
// future, which values will print from the goroutine
// are undetermined

// if we instead pass in the objects as shown below
// the function scope will create a deep copy of the
// items within the new scope, giving us predictable
// outcomes
func intro02memoryB() {
	var wg sync.WaitGroup // this is a WaitGroup
	for i, item := range []string{"a", "b", "c", "d"} {
		wg.Add(1)
		go func(item string, i int) {
			defer wg.Done()
			fmt.Printf("| $%s : %d |", item, i)
		}(item, i)
	}
	wg.Wait()
}

// go routines are given a few kb to start, if it's not
// enough, runtime will grow or reduce it as necessary.
// This allows many goroutines to exist, and given
// that it takes only 3-5 inexpensive instructions per
// function call, it is fairly inexpensive and practical
// to instantiate hundreds of thousands of goroutines
// without straining system resources.
func intro02memoryC() {
	var wg sync.WaitGroup
	// change this   v number to see what happens
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("[%d]", i)
		}(i)
	}
	wg.Wait()
}

// However note that the garbage collector does not have
// the responsibility of closing routines.
// if a routine has been abandoned, it will exist until
// the process exits. This is referred to as a "hang"
/*
func hangfunc() {
	go func() {
		for{
			// hangs forever
		}
	}()
}
*/
