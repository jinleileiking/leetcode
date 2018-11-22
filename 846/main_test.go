package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestGetW(t *testing.T) {

	spew.Dump(getW([]int{4, 3, 3, 2, 2, 1}, 3))
}

func TestGetW1(t *testing.T) {

	spew.Dump(getW([]int{3, 2, 1}, 3))
}

func TestMain(t *testing.T) {

	spew.Dump(isNStraightHand([]int{1, 2, 2, 3, 3, 4}, 3))
}

func TestMain1(t *testing.T) {

	spew.Dump(isNStraightHand([]int{1, 7, 2, 3, 3, 4}, 3))
}

func TestMain2(t *testing.T) {

	spew.Dump(isNStraightHand([]int{1, 2, 3}, 1))
}
