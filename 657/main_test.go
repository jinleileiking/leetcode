package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestMain(t *testing.T) {
	// spew.Dump(judgeCircle("LL"))
	// spew.Dump(judgeCircle("UD"))
	spew.Dump(judgeCircle("RLUURDDDLU"))

	// D D D R R R L LL U
	//spew.Dump(judgeCircle("DURDLDRRLL"))

	// U

	// UUULLL

	// DLDLDDRRRDDRUDUDRDLDRRUDLUDDRUDULULDDDLUULDLLURR
	// spew.Dump(judgeCircle("UUDDLULDULRDRRRULLLRRUDUURUULLDLLLRDLDLLRDRLDRRRDDRUDUDLRDLDRRUDLRLUDDRULLRRDLRLRLRULULDDDLUULDLLURR"))

}
