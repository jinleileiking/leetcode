package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	max := 0
	for p, _ := range s {
		max = find(s[p:], max)
	}

	if max == 0 {
		return len(s)
	}
	return max

}

func find(s string, curMax int) int {
	var stat map[rune]bool
	var max int
	stat = make(map[rune]bool)

	var kfinal int
	for k, v := range s {
		kfinal = k
		if _, exists := stat[v]; exists {
			// repeated
			if k > max {
				max = k
			}
			break
		} else {
			stat[v] = true
		}
	}

	if kfinal == len(s)-1 {
		if len(s) > curMax {
			return len(s)
		}
	}

	if max > curMax {
		return max
	}

	return curMax
}
