package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// spew.Dump(maxValue("123", 4))
	// spew.Dump(maxValue("1345", 4))
	// spew.Dump(maxValue("13475", 4))
	spew.Dump(maxValue("-13475", 4))

}

func maxValue(n string, x int) string {
	isNeg := false

	digNum := len(n) + 1
	digs := []int{}

	for _, c := range n {
		if c == '-' {
			isNeg = true
		} else {
			if i, err := strconv.Atoi(string(c)); err == nil {
				digs = append(digs, i)
			}
		}
	}

	digs = append(digs, x)

	if isNeg {
		digNum = digNum - 1
		return fmt.Sprintf("%d", findMin(digs, digNum))
	} else {
		return fmt.Sprintf("%d", findMax(digs, digNum))
	}

}

func findMax(digs []int, digNum int) int {
	ret := 0
	l := digNum
	left := digs
	max := 0
	for i := 0; i < l; i++ {
		max, left = findMaxDig(left)
		ret = ret + max*int(math.Pow(10, float64(l-i-1)))
	}

	return ret
}

func findMaxDig(digs []int) (max int, left []int) {
	maxP := 0
	for p, v := range digs {
		if v > max {
			max = v
			maxP = p
		}
	}

	left = append(digs[:maxP], digs[maxP+1:]...)

	return max, left
}

func findMin(digs []int, digNum int) int {
	ret := 0
	l := digNum
	left := digs
	min := 0
	for i := 0; i < l; i++ {
		min, left = findMinDig(left)
		spew.Dump(min, left)
		ret = ret + min*int(math.Pow(10, float64(l-i-1)))
		spew.Dump(ret)
	}

	return -ret
}

func findMinDig(digs []int) (min int, left []int) {
	min = 10
	minP := 0
	for p, v := range digs {
		if v < min {
			min = v
			minP = p
		}
	}

	// spew.Dump(minP)
	left = append(digs[:minP], digs[minP+1:]...)

	return min, left
}
