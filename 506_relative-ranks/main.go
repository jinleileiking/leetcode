package main

import (
	"sort"
	"strconv"
)

func findRelativeRanks(nums []int) []string {
	var src = make([]int, len(nums))
	var ret []string = make([]string, len(nums))

	copy(src, nums)
	sort.Slice(nums, func(i int, j int) bool { return nums[i] > nums[j] })

	m := make(map[int]int)

	// 9, 8 ,7 , 6
	// r 0  1
	// v 9, 8

	for r, v := range nums {
		m[v] = r
	}

	for k, v := range src {
		if r, ok := m[v]; ok {
			if r == 0 {
				ret[k] = "Gold Medal"
			} else if r == 1 {
				ret[k] = "Silver Medal"
			} else if r == 2 {
				ret[k] = "Bronze Medal"
			} else {
				ret[k] = strconv.Itoa(r + 1)
			}
		}
	}
	return ret
}
