package main

import (
	"reflect"
	"sort"
)

func threeSum(nums []int) [][]int {

	ret := [][]int{}
	// foundKeys := []int{}

	m := make(map[int][][][]int)
	for k1, v1 := range nums {
		for k2 := k1 + 1; k2 < len(nums); k2++ {
			v2 := nums[k2]

			if _, ok := m[v2]; ok {
				//found
				// spew.Dump(m, v2)
				retA := [][][]int{}
				ret, retA = insert(ret, v2, m[v2], k2)

				if len(retA) != 0 {
					m[v2] = retA
				} else {
					delete(m, v2)
				}
			}

			// not found, save map   5 --> -1, -4,  -2, -3
			toFind := -v1 - v2

			// if 2 == toFind {
			// 	spew.Dump(v1, v2, k1, k2)
			// }

			aInner := [][]int{
				[]int{v1, k1},
				[]int{v2, k2},
			}

			m[toFind] = append(m[toFind], aInner)
		}
	}

	return ret
}
func insert(ret [][]int, v int, as [][][]int, p int) ([][]int, [][][]int) {

	retA := [][][]int{}
	// a ---> [2  3],  [4 5]     2,4,-6
	for _, a := range as {

		pass := false
		t := []int{v}
		for _, item := range a {
			t = append(t, item[0])
			if item[1] == p {
				pass = true
				break
			}
		}

		if pass {
			retA = append(retA, a)
			continue
		}
		// this a should not set to int, and should not delete

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

	return ret, retA
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
