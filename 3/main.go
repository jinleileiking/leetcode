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
		// spew.Dump(s[p:], max)
		max = find(s[p:], max)
	}

	if max == 0 {
		return len(s)
	}
	return max

}

func find(s string, curMax int) int {
	var stat map[rune]bool
	stat = make(map[rune]bool)

	for k, v := range s {
		// spew.Dump(k)
		if _, exists := stat[v]; exists {
			// repeated
			if k >= curMax {
				return k
			}
			return curMax
		} else {
			stat[v] = true
		}
	}

	if len(s) > curMax {
		return len(s)
	}

	return curMax
}
