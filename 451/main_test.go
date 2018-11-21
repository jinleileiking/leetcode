package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestMain(t *testing.T) {

	spew.Dump(frequencySort("aaaaAAAb"))
}
