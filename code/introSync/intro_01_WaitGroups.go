package introSync

import "sync"
import util "GoPatterns/code/util"

// There's a problem with our previous code :
/* 0
// Basic building blocks of go concurrency
func intro00GoRoutines() {
	for i := 0; i < 2; i++ {
		go intro00doTheseThings()
	}
	// continue to do other things
	for i := 0; i < 5; i++ {
		fmt.Print("|Other Thing|")
	}
	time.Sleep(1 * time.Millisecond)
}

func intro00doTheseThings() {
	fmt.Print("|The Thing|")
}
*/
/* 1
// and that is that it is not guaranteed to run at all.
// coroutines are, for our intents and purposes here, independent unless
// defined to be otherwise except in one way:
// If the main thread finishes, the goroutine will end as well, whether it has
// completed its job or not.
// We can solve for this in a number of ways, using a number of patterns, but
// here we are only going to introduce one of the primitives that can be used
// to do this
*/

func intro01WaitGroupsA() {
	var wg sync.WaitGroup // this is a WaitGroup
	wg.Add(1)
	go util.AThing{ThingType: "A Thing", Thread: 0}.DoWithWG(&wg)
	wg.Wait()
	// if we comment out this wait group, there is no guarantee that
	// A Thing will finish before this function scope closes.
}

// Now we can do this with a large number of routines without much of a problem
func intro01WaitGroupsB() {
	var wg sync.WaitGroup // this is a WaitGroup
	for i := 0; i < 25; i++ {
		wg.Add(1)
		go util.AThing{ThingType: "A Thing", Thread: i}.DoWithWG(&wg)
	}
	wg.Wait()
}
