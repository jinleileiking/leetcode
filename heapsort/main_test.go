package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestHeapsort(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6}

	spew.Dump(Heapsort(a))
}

func TestHeapsort1(t *testing.T) {
	a := []int{2, 3, 4, 5, 6}

	spew.Dump(Heapsort(a))
}

func TestChange(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6}

	change(&a, 6)

	spew.Dump(a)
}

func TestChange1(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5}

	change(&a, 5)

	spew.Dump(a)
}

func TestChange2(t *testing.T) {
	a := []int{5, 4, 0, 3, 1, 2}

	change(&a, 5)

	spew.Dump(a)
}

func TestHeapify(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6}

	Heapify(a)

	spew.Dump(a)
}

func TestHeapify1(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5}

	Heapify(a)

	spew.Dump(a)
}
