package intro

import (
	"fmt"
	"sync"
	"time"
)

// So what do we have so far? we've looked at waitgroups, got a basic
// understanding of routines and a bit about their memory.. We learned about
// mutex locking and how to use RW locks to prevent unnecessary locking..
// so let's move on and suggest a scenario where we have to do a little more
// than wait for two routines to finish... instead, we have two coroutines, one
// is going to add something to an int, and then the other is going to put the
// int on a string and send it somewhere. coroutine b can't do its thing
// until a did its thing. How do we signal to b that the thing happened and that
// it can now do its job?
// (hint for those who read ahead: does not have to be a channel)
// we could pass in some other variable for it to listen to?

// Intro05primitiveA .
func Intro05primitiveA(ct int) {
	Increment := func(a *int, b *bool, wg *sync.WaitGroup) {
		time.Sleep(1 * time.Nanosecond)
		defer wg.Done()
		*a++
		*b = true
		//fmt.Println("done incrementing")
	}
	Write := func(a *int, b *bool, wg *sync.WaitGroup) {
		defer wg.Done()
		//fmt.Println("entering loop")
		for !*b {
			time.Sleep(1 * time.Millisecond)
		} // endless loop, until b == true.
		_ = fmt.Sprintf("[%d]", *a)
		//fmt.Println(s)
	}
	var wg sync.WaitGroup
	for i := 0; i < ct; i++ {
		a, b := 5, false
		wg.Add(2)
		go Increment(&a, &b, &wg)
		go Write(&a, &b, &wg)
	}
	wg.Wait()
}

// We're gonna look at two tables here, one
/*time to completion
second	     1,054,923,800.00
milisecond	 311,167,700.00
micro	       420,873,900.00
nanosecond	 604,388,300.00
none	       705,112,900.00
*/
// Now this is interesing, because we'd think that removing a sleep from the
// middle of our Write function would increase speed, but it doesn't, in fact,
// all the way up to a second of time, adding a sleep actually decreases runtime.
// not only will this function also often hang, but it will also use a significant
// amount of CPU
// in one case, where it managed to complete, the cpu profile showed it spent
// more than 25% of its time running just sitting in for{}
// even if nothing happens in that loop, the machine is still looping.

// There is no comparison between the time it takes the above to complete vs
// the use of the one below:

// on my machine I get 558,505,700 a. vs 997,800 for b. (nanoseconds), and b
// is far more consistent, whereas a can be anywhere between 558 to 900 million
// nanoseconds.

// Intro05primitiveB .
func Intro05primitiveB(ct int) {
	Increment := func(a *int, bWg *sync.WaitGroup, wg *sync.WaitGroup) {
		time.Sleep(1 * time.Nanosecond)
		defer wg.Done()
		defer bWg.Done()
		*a++
		//fmt.Println("done incrementing")
	}
	Write := func(a *int, bWg *sync.WaitGroup, wg *sync.WaitGroup) {
		defer wg.Done()
		//fmt.Println("entering loop")
		bWg.Wait()
		_ = fmt.Sprintf("[%d]", *a)
		//fmt.Println(s)
	}
	var wg sync.WaitGroup
	for i := 0; i < ct; i++ {
		a, bWg := 5, sync.WaitGroup{}
		wg.Add(2)
		bWg.Add(1)
		go Increment(&a, &bWg, &wg)
		go Write(&a, &bWg, &wg)
	}
	wg.Wait()
}

// while this works, it's kind of hacky, and waitgroups aren't meant to be used
// this way--not to say that they can't.
// to avoid being hacky, we can instead use sync.condition.
// now, the following is considered "unsafe" and can hang, I will produce a
// better contrived example in the future.

// That being said, using condition is considered best practise over waitgroups
// in many cases. We should use waitgroups when we have to wait for blocks of
// routines to finish in order to continue. However, when we want to communicate
// between ROUTINES, in order to allow certain groups to continue condition can
// be far more advantageous

// In this example, we feed a condition, which takes a mutex pointer as an arg.
// When we use cond.wait() it actually suspends the goroutine until a "signal"
// is sent. Behind the curtain, there is a FIFO (first in first out) lsit of
// routines waiting for their condition to signal. When signal is fired, what it
// does is go to the routine that has been waiting the longest and notifies it
// to unwait.

// Intro05Cond .
func Intro05Cond(ct int) {
	Increment := func(a *int, cond *sync.Cond, wg *sync.WaitGroup) {
		time.Sleep(1 * time.Nanosecond)
		defer wg.Done()
		cond.L.Lock()
		*a++
		cond.L.Unlock()
		//fmt.Printf("[%d>%d][SIG-OUT]", x, *a)
		cond.Signal()
		//fmt.Println("done incrementing")
	}
	Write := func(a *int, cond *sync.Cond, wg *sync.WaitGroup) {
		defer wg.Done()
		//fmt.Println("entering loop")
		cond.L.Lock() // entering critical section
		//fmt.Print("[SIG-WAIT]")
		cond.Wait()     // this will suspend the routine until a signal is received
		cond.L.Unlock() // exit c.s.
		_ = fmt.Sprintf("[%d]", *a)
		//fmt.Println(s)
	}
	var wg sync.WaitGroup
	for i := 0; i < ct; i++ {
		a, cond := 5, sync.NewCond(&sync.Mutex{})
		wg.Add(2)
		go Write(&a, cond, &wg)
		go Increment(&a, cond, &wg)
	}
	wg.Wait()
}

// Sync.cond has a second method which is worth mentioning, but I do not have
// the time to upload an example of at this time.. that method is "broadcast".
// Where signal sends to the routine that has been waiting the longest..
// broadcast sends to EVERY routine on the FiFo list of routines.
