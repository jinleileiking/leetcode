package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestMain(t *testing.T) {
	src := []int{1, 2, 3}
	spew.Dump(plusOne(src))

}
func TestMain1(t *testing.T) {
	src := []int{4, 3, 2, 1}
	spew.Dump(plusOne(src))

}

func TestMain2(t *testing.T) {
	src := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	spew.Dump(plusOne(src))

}
