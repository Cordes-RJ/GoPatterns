package introSync

import (
	"fmt"
	"sync"
)

// continuing from mutexes and memory, it should be
// noted that the mutex is stopping access, each time
// it locks, from all other coroutines, this means
// that each co-routine has to wait until it's done
// in order to continue.

// This will cause bottlenecks in your program. These
// moments where everything else is blocked from access
// are sometimes referred to as ""critical sections" and
// time within them should be limited.
// if there's no option but to use mutex, the nature of the
// access should be considered. Does the routine need to
// write over the resource? or does it just need to read it?
// does it ABSOLUTELY need the MOST CURRENT version when it
// does?
// There is a special mutex, an RWMutex, to use when you want
// to lock types of access to the resource instead of just
// locking the resource.

// There's a lot of code below, but I'll try to summarise.
// funcA is using a rwmutex, whereas b is using a traditional
// mutex. We "go" create 1000 incrementers, decrementers and
// observers. The incrementers add 1 to each item in an array,
// the decrementers minus 1, and the observers simply read out
// from the array.
// if we lock access every time, even when we only
// need to read, (which could be done by many routines at once)
// this can greatly slow down our program.
// On my machine, the test differed greatly, run 100 times:
// 966415700
// vs
//   7979200
//
//
func intro04makeIncrementer(array []int, wg *sync.WaitGroup, L sync.Locker) {
	defer wg.Done()
	L.Lock()
	defer L.Unlock()
	for i := range array {
		array[i]++
	}
}
func intro04makeDecrementer(array []int, wg *sync.WaitGroup, L sync.Locker) {
	defer wg.Done()
	L.Lock()
	defer L.Unlock()
	for i := range array {
		array[i]--
	}
}

func intro04makeObserver(array []int, wg *sync.WaitGroup, L sync.Locker) {
	defer wg.Done()
	L.Lock()
	defer L.Unlock()
	for i := range array {
		_ = array[i]
	}
}

type intro04makeThing func(array []int, wg *sync.WaitGroup, L sync.Locker)

func intro04makeCt(makeThing intro04makeThing, array []int, wg *sync.WaitGroup, L sync.Locker, ct int) {
	for i := 0; i < ct; i++ {
		wg.Add(1)
		go makeThing(array, wg, L)
	}
}

func intro04rwMutexA() {
	var array []int
	for i := 0; i < 1000; i++ {
		array = append(array, 0)
	}
	var rw = &sync.RWMutex{}
	var m = &sync.Mutex{}
	var wg sync.WaitGroup
	go intro04makeCt(intro04makeIncrementer, array, &wg, m, 1000)
	go intro04makeCt(intro04makeDecrementer, array, &wg, m, 1000)
	go intro04makeCt(intro04makeObserver, array, &wg, rw, 1000)
	wg.Wait()
	if array[0] != 0 {
		fmt.Println("err")
	}
}

func intro04rwMutexB() {
	var array []int
	for i := 0; i < 1000; i++ {
		array = append(array, 0)
	}
	var m = &sync.Mutex{}
	var wg sync.WaitGroup
	go intro04makeCt(intro04makeIncrementer, array, &wg, m, 1000)
	go intro04makeCt(intro04makeDecrementer, array, &wg, m, 1000)
	go intro04makeCt(intro04makeObserver, array, &wg, m, 1000)
	wg.Wait()
	if array[0] != 0 {
		fmt.Println("err")
	}
}
