package introChannels

import (
	"fmt"
	"sync"
)

// Sync is incredibly useful whgen we're dealing with performance critical
// sections with shared memory for Read Write access, but not so great when
// we're trying to transfer ownership over data or coordinate multiple pieces
// of logic. For these things, we have the synchronisation primitive that golang
// is famous for, the channel. Note though, that channels are not a panacea,
// channels have many uses but they can hurt more than help if used incorrectly.

// We're going to start with basic syntax

// Intro07ChannelSyntax1 is a basic introduction to channel syntax
func Intro07ChannelSyntax1() {
	fmt.Println("testing")
	var pipe chan int     // [1]
	pipe = make(chan int) // [2]
	//
	// pipeSendOnly := make(chan<- int)    // [3]
	// pipeReceiveOnly := make(<-chan int) // [4]
	var wg sync.WaitGroup // [5]
	Receiver := func(C <-chan int) {
		defer wg.Done()
		fmt.Printf("[<-%d]", <-C)
	} // [6]
	Sender := func(C chan<- int, v int) {
		defer wg.Done()
		fmt.Printf("[%d<-]", v)
		C <- v
	} // [7]
	for i := 0; i < 10; i++ {
		wg.Add(2)
		pipe = make(chan int)
		go Receiver(pipe)
		go Sender(pipe, i)
	} // [8]
	wg.Wait()
}

// [1] Declare the channel and its type
// [2] Instantiate the channel with make()
// [3] pipeSendOnly is declared and made in one line here,
//     and this channel will only allow things to be sent through it.
// [4] pipeReceiveOnly is declared to only receive
// [3,4] Why use send/receive only channels? what good is that?
//       Well, this is a good opportunity to look at an interesting feature of
//       of channels.. Bidirectional channels can change type implicitly, so
//       read and send are normally used as return and argument types.
// [5] Note the return of the sync package. Use of sync and channel primitives
//     are note mutually exclusive
// [6] We create a function which takes a READ ONLY channel as an argument,
//     it reads in the value and prints it
// [7] We create a function which takes a SEND ONLY channel as an argument,
//     it takes an integer as an argument as well, and will send it through the
//     channel and print to document that it has sent successfully.
// [8,3,4] In the loop here, we show that the bidirectional channel can be used
//     as a readonly AND sendonly argument, as it will implicitly change the
//     the type of the bidirectional channel in the new scope.
