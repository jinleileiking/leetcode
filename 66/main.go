package main

import (
	"math"

	"github.com/davecgh/go-spew/spew"
)

func pow(src, up int) uint64 {
	return uint64(math.Pow(float64(src), float64(up)))
}

func plusOne(digits []int) []int {
	var src uint64
	for i := len(digits) - 1; i >= 0; i-- {
		src = src + uint64(digits[len(digits)-i-1])*pow(10, i)
	}

	spew.Dump(src)
	src = src + 1

	ret := []int{}
	var i int
	for {
		if src/pow(10, i) != 0 {
			ret = append(ret, int(src/pow(10, i)%10))
			i++
		} else {
			break
		}
	}

	// spew.Dump(ret)
	ret1 := make([]int, len(ret))
	for k, _ := range ret {
		ret1[k] = ret[len(ret)-1-k]
		ret1[len(ret)-1-k] = ret[k]
		if k > len(ret)/2 {
			break
		}
	}
	// spew.Dump(ret)
	return ret1
}
