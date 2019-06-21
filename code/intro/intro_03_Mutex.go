package intro

import (
	"fmt"
	"strconv"
	"sync"
)

// continuing on the topic of memory, is the shared usage
// by co-routines of resources found in other co-routines
// or within the main thread

// Let's increment a counter in main, by using external
// routines.
func intro03mutexA() {
	var counter int
	var lock sync.Mutex
	// we can use mutex to sync access
	addToCtr := func(wg *sync.WaitGroup) {
		defer wg.Done()
		lock.Lock()
		defer lock.Unlock()
		// remember to defer an unlock, failing to include This
		// may cause your program to "deadlock"
		from := counter
		counter++
		fmt.Printf("| [%d]->[%d]", from, counter)
	}
	remFromCtr := func(wg *sync.WaitGroup) {
		defer wg.Done()
		lock.Lock()
		defer lock.Unlock()
		from := counter
		counter--
		fmt.Printf("| [%d]->[%d] |", from, counter)
	}
	// now add a wait group
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		if (i % 2) == 0 {
			wg.Add(1)
			go addToCtr(&wg)
		} else {
			wg.Add(1)
			go remFromCtr(&wg)
		}
	}
	wg.Wait()
	fmt.Println(strconv.Itoa(counter))
}

// if we do this without a sync lock
// We will no longer get the tit for tat result of 0
// instead we can end up with random results, in my test
// I received -8
func intro03mutexB() {
	var counter int
	//var lock sync.Mutex
	// we can use mutex to sync access
	addToCtr := func(wg *sync.WaitGroup) {
		defer wg.Done()
		//lock.Lock()
		//defer lock.Unlock()
		//from := counter
		counter++
		//fmt.Printf("| [%d]->[%d]", from, counter)
	}
	remFromCtr := func(wg *sync.WaitGroup) {
		defer wg.Done()
		//lock.Lock()
		//defer lock.Unlock()
		//from := counter
		counter--
		//fmt.Printf("| [%d]->[%d] |", from, counter)
	}
	// now add a wait group
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		if (i % 2) == 0 {
			wg.Add(1)
			go addToCtr(&wg)
		} else {
			wg.Add(1)
			go remFromCtr(&wg)
		}
	}
	wg.Wait()
	fmt.Println(strconv.Itoa(counter))
}
