package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestMedianOnce(t *testing.T) {

	// spew.Dump(findMedianFloat([]float64{1.0, 2.0, 3.0}))
	// spew.Dump(findMedianFloat([]float64{1.0, 2.0, 3.0, 4.0, 5.0}))
	spew.Dump(findMedianFloat([]float64{2.0, 2.0}))
}

func TestMedian1(t *testing.T) {

	spew.Dump(findMedian([]int{1, 2, 3}))
	spew.Dump(findMedian([]int{1, 2, 3, 4, 5}))
}
