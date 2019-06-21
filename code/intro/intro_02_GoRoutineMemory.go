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
// but instead, we get d:3 d:3 d:3...
// You may also get an error on the line, sugges

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
