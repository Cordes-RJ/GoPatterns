package intro

import (
	"fmt"
	"time"
)

// Basic building blocks of go concurrency
func intro00GoRoutines() {
	for i := 0; i < 2; i++ {
		go intro00doTheseThings()
	}
	// continue to do other things
	for i := 0; i < 5; i++ {
		fmt.Printf("|[%d]Other Thing|", i)
	}
	time.Sleep(1 * time.Millisecond)
}

func intro00doTheseThings() {
	fmt.Printf("|The Thing|")
}

// Goroutines are unique to go, though other languages have similar primitives
// They're no os threads, and they're not exeactly green threads, they're
// called co routines. They're deeply integrated with go's runtime. Goroutines
// don't define their own suspension or re-entry from a fork
// you define re-entry-- but they don't even have to rejoin.
// go will suspend them if they block and will resume them if they become
// unblocked.
