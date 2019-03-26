package main

func removeKdigits(num string, k int) string {
	var last rune
	ret := ""
	deleted := 0
	// found1st0 := false
	for _, v := range num {

		// first
		if byte(last) == 0 {

			// 0000xxxxx
			if v == '0' {
				continue
			}

			last = v
			ret += string(last)
			continue
		}

		if deleted == k {
			ret += string(v)
			continue
		}

		// 10200 -> 200
		// [1] --> []
		if v == '0' && k != deleted {
			if k < len(ret) {
				// k >=  len ret
				// 112340       3
				// 110         deleted = 3
				ret = ret[:k]
				ret += string(v)
				deleted = k
			} else {
				// k >=  len ret
				// 112340       8
				// 00000         deleted = 5

				ret = ""
				deleted += k - len(ret)
				last = 0

				// spew.Dump(deleted, k)
			}

			continue
		}

		if v < last {

			// 5 3
			// 3

			// xxxx53  xxxx5 0 -4  0 - len -2
			// xxxx3
			if len(ret) == 1 {
				ret = string(v)
				// } else if len(ret) == 2 {
				// 	// x 5  3
				// 	ret = ret[:0]
				// 	ret += string(v)
			} else {
				// spew.Dump(ret)
				ret = ret[:len(ret)-2]
				ret += string(v)
			}

			last = v
			deleted++
		} else if v > last {
			// xxxx35
			// xxxx3

			// corner case  002
			// if last == '0' {
			// 	last = v
			// 	continue
			// }

			deleted++
		} else {
			ret += string(v)

			last = v
		}
	}

	if ret == "" {
		return "0"
	}

	if k == len(num) {
		return "0"
	}

	if deleted < k {
		// 1111111, 3
		// 0123
		// 0: 7 - 4
		// spew.Dump(deleted)
		return ret[:len(ret)-(k-deleted)]
	}
	return ret

}
