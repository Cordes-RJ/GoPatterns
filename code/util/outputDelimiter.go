package util

import "fmt"

// OutputDelimiter is a structure which facilitates the delimiting of console
// output by named scope
type OutputDelimiter struct {
	name string
}

// DelimitOutput marks the start of scope and returns an object which can be
// used to exit the scope
func DelimitOutput(funcName string) OutputDelimiter {
	od := OutputDelimiter{funcName}
	od.start()
	return od
}

// start marks entrance of scope
func (od OutputDelimiter) start() {
	fmt.Printf("<%s>\n", od.name)
}

// End marks exit of scope
func (od OutputDelimiter) End() {
	fmt.Printf("\n</%s>\n", od.name)
}
