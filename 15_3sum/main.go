package main

import (
	"reflect"
	"sort"

	"github.com/davecgh/go-spew/spew"
)

func threeSum(nums []int) [][]int {

	// ret := [][]int{}
	foundKeys := []int{}

	m := make(map[int][][]int)
	for k1, v1 := range nums {
		for k2 := k1 + 1; k2 < len(nums); k2++ {
			v2 := nums[k2]

			if _, ok := m[v2]; ok {
				//found
				// ret = insert(ret, v2, m[v2])
				foundKeys = append(foundKeys, v2)
			}

			// not found, save map   5 --> -1, -4,  -2, -3
			toFind := -v1 - v2
			m[toFind] = append(m[toFind], []int{v1, v2})
		}
	}

	spew.Dump(m)
	spew.Dump(foundKeys)

	return genRet(m, foundKeys)
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

// func insert(ret [][]int, v []int) [][]int {

// 	// sort

// 	a := []int{k1, k2, k3}
// 	sort.Slice(a, func(i int, j int) bool { return a[i] < a[j] })

// 	for _, v := range ret {
// 		if reflect.DeepEqual(a, v) {
// 			return ret
// 		}
// 	}

// 	ret = append(ret, a)
// 	return ret
// }
