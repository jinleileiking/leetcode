package main

func checkSubarraySum(nums []int, k int) bool {

	sum := 0

	// mod -> pos
	m := make(map[int]int)

	for k1, v := range nums {

		sum = sum + v
		mod := sum % k
		if _, ok := m[mod]; ok {
			//found same sum
			return true
		}

		// not found
		m[mod] = k1

	}

	return false
}
