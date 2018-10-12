package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestStr(t *testing.T) {
	// spew.Dump(findStr("aabzZ1234"))
	// spew.Dump(findStr("1234"))
	// spew.Dump(findStr("aa1234"))
	// spew.Dump(findStr("zzzz"))
	// spew.Dump(findStr("aabzZ2"))
	spew.Dump(findStr("a2345678999999999999999"))
}

func TestInt(t *testing.T) {
	// spew.Dump(findInt("aabzZ1234"))
	// spew.Dump(findInt("1234"))
	// spew.Dump(findInt("1234aaa"))
	// spew.Dump(findInt("2"))
	spew.Dump(findInt("2345678999999999999999"))
}

func TestMain(t *testing.T) {
	spew.Dump(decodeAtIndex("aabzZ2", 5))
	// spew.Dump(decodeAtIndex("a2345678999999999999999", 1))

}
