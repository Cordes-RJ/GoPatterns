package intro

import (
	"fmt"
	"testing"
)

func TestIntro00(t *testing.T) {
	name := "intro_00_Routines"
	fmt.Printf("<%s>\n", name)
	defer fmt.Printf("</%s>\n", name)
	intro00GoRoutines()
	x := false
	if x {
		t.Error("error")
	}
}
