package utility

import (
	"fmt"
	"strconv"
)

// Thing is a thing which a function can do and finish (done)
type Thing interface {
	Do()
	Done()
	MakePacket() Packet
}

// AThing is a simple thing which has a type and a thread
type AThing struct {
	ThingType string
	Thread    int
}

// AThing's Do() just prints its thingtyp, thread and state
func (thing AThing) Do() {
	fmt.Printf("|%s|%d|Working...|", thing.ThingType, thing.Thread)
}

// AThing's Done() just prints its thingtype, thread and state
func (thing AThing) Done() {
	fmt.Printf("|%s|%d|Done|", thing.ThingType, thing.Thread)
}

// MakePacket simply returns a packet
func (thing AThing) MakePacket() Packet {
	return Packet{thing.ThingType, strconv.Itoa(thing.Thread)}
}
