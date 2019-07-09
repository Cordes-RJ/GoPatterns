package introChannels

import "testing"
import "GoPatterns/code/util"

func TestIntro07ChannelSyntax1(t *testing.T) {
	d := util.DelimitOutput("Intro07ChannelSyntax1")
	defer d.End()
	Intro07ChannelSyntax1()
	if false {
		t.Error()
	}
}

func TestRangeOverChan08(t *testing.T) {
	d := util.DelimitOutput("RangeOverChan08")
	defer d.End()
	RangeOverChan08()
	if false {
		t.Error()
	}
}
