package main

func checkSubarraySum(nums []int, k int) bool {

	sum := 0

	// mod -> pos
	m := make(map[int]int)

	for k1, v := range nums {

		// corner check
		if len(nums) == 1 {
			return false
		}

		if k == 0 && v != 0 {
			return false
		}

		// corner check
		if k == 0 && v == 0 {
			if k1 == len(nums)-1 {
				return true
			}
			continue
		}

		sum = sum + v

		// if sum%k == 0 {
		// 	return true
		// }

		mod := sum % k
		if _, ok := m[mod]; ok {
			//found same sum
			return true
		}

		// not found
		m[mod] = k1

	}

	//corner check

	// only one mods
	if p, ok := m[0]; ok {
		if p == 0 {
			return false
		}
		return true
	}

	return false
}
