package utility

import (
	"fmt"
	"strconv"
	"sync"
)

// Thing is a thing which a function can do and finish (done)
type Thing interface {
	DoWithWG()
	Done()
	MakePacket() Packet
}

// AThing is a simple thing which has a type and a thread
type AThing struct {
	ThingType string
	Thread    int
}

// DoWithWG just prints its thingtyp, thread and state
// but does so with a defer waitgroup
func (thing AThing) DoWithWG(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("|$%s[%d]:W|", thing.ThingType, thing.Thread)
	thing.Done()
}

// Done just prints its thingtype, thread and state
func (thing AThing) Done() {
	fmt.Printf("|$%s[%d]:D|", thing.ThingType, thing.Thread)
}

// MakePacket simply returns a packet
func (thing AThing) MakePacket() Packet {
	return Packet{thing.ThingType, strconv.Itoa(thing.Thread)}
}

// MakeAThing makes AThing
func MakeAThing(typeofThing string, thread int) AThing {
	return AThing{typeofThing, thread}
}
