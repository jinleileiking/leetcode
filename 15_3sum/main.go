package main

import (
	"reflect"
	"sort"

	"github.com/davecgh/go-spew/spew"
)

func threeSum(nums []int) [][]int {

	ret := [][]int{}
	// foundKeys := []int{}

	m := make(map[int][][]int)
	for k1, v1 := range nums {
		for k2 := k1 + 1; k2 < len(nums); k2++ {
			v2 := nums[k2]

			if _, ok := m[v2]; ok {
				//found
				ret = insert(ret, v2, m[v2])
				//del

				spew.Dump(m, v2)
				delete(m, v2)
			}

			// not found, save map   5 --> -1, -4,  -2, -3
			toFind := -v1 - v2
			m[toFind] = append(m[toFind], []int{v1, v2})
		}
	}

	return ret
}
func insert(ret [][]int, v int, as [][]int) [][]int {

	// sort

	for _, a := range as {
		t := []int{v, a[0], a[1]}
		sort.Slice(t, func(i int, j int) bool { return t[i] < t[j] })

		found := false
		for _, v := range ret {
			if reflect.DeepEqual(t, v) {
				found = true
				break
			}
		}

		if !found {
			ret = append(ret, t)
		}
	}

	return ret
}

func genRet(m map[int][][]int, foundKeys []int) [][]int {

	ret := [][]int{}

	for _, v := range foundKeys {
		for _, a := range m[v] {
			t := []int{v, a[0], a[1]}
			sort.Slice(t, func(i int, j int) bool { return t[i] < t[j] })

			found := false
			for _, v := range ret {
				//found in ret
				if reflect.DeepEqual(t, v) {
					found = true
					break
				}
			}

			if !found {
				ret = append(ret, t)
			}
		}
	}

	return ret
}
