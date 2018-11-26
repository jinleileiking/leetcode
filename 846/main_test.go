package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

// func TestGetW(t *testing.T) {

// 	spew.Dump(getW([]int{4, 3, 3, 2, 2, 1}, 3))
// }

// func TestGetW1(t *testing.T) {

// 	spew.Dump(getW([]int{3, 2, 1}, 3))
// }

func TestMain(t *testing.T) {

	spew.Dump(isNStraightHand([]int{1, 2, 2, 3, 3, 4}, 3))
}

func TestMain1(t *testing.T) {

	spew.Dump(isNStraightHand([]int{1, 7, 2, 3, 3, 4}, 3))
}

func TestMain2(t *testing.T) {

	spew.Dump(isNStraightHand([]int{1, 2, 3}, 1))
}
func TestMain3(t *testing.T) {

	spew.Dump(isNStraightHand([]int{1, 2, 2, 3, 3, 3, 4, 4, 5}, 2))
}

func TestHasw(t *testing.T) {

	spew.Dump(hasW([]int{1, 2, 3, 4}, 3, 4))
}

func TestHasw1(t *testing.T) {

	spew.Dump(hasW([]int{3, 3, 3, 1}, 2, 3))
}

func TestHasw2(t *testing.T) {

	spew.Dump(hasW([]int{3, 2, 1}, 2, 3))
}

func TestCutw(t *testing.T) {

	spew.Dump(cutW([]int{1, 2, 2, 1, 3, 4}, 3, 4))
}

func TestCutw1(t *testing.T) {

	spew.Dump(cutW([]int{4, 3, 3, 3, 2, 2, 1}, 2, 4))
}

func TestCutw3(t *testing.T) {

	spew.Dump(cutW([]int{3, 2, 1}, 2, 3))
}
