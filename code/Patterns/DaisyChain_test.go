package GoPattern

import (
	"GoPatterns/code/util"
	"testing"
)

func TestDaisyChain(t *testing.T) {
	d := util.DelimitOutput("DaisyChain")
	defer d.End()
	DaisyChain()
	if false {
		t.Error()
	}
}
