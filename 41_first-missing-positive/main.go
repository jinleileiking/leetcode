package main

func findMax(nums []int) int {
	var max int

	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func firstMissingPositive(nums []int) int {

	max := findMax(nums)

	space := make([]int, max+1)

	for _, v := range nums {
		if v <= 0 {
			continue
		}
		space[v] = 1
	}

	for k, v := range space {
		if k == 0 {
			continue
		}
		if v == 0 {
			return k
		}
	}

	return len(nums) + 1
}
