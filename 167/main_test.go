package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestMain(t *testing.T) {

	spew.Dump(twoSum([]int{1, 2, 3, 4}, 5))
}
