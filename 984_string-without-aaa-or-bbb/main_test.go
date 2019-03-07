package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestMain(t *testing.T) {

	spew.Dump(strWithout3a3b(1, 2))
	spew.Dump(strWithout3a3b(1, 1))
	spew.Dump(strWithout3a3b(4, 1))
	// assert.Equal(t, "abb", strWithout3a3b(1, 2))
	// assert.Equal(t, "aabaa", strWithout3a3b(4, 1))
}
