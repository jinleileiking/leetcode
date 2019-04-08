package main

func minFlipsMonoIncr(S string) int {

	minFlip := -1
	cnt0Total := 0

	last := 0
	allSame := true
	for _, v := range S {
		if v == '0' {
			cnt0Total++
		}

		if last == 0 {
			last = int(v)
		}

		if last != int(v) {
			allSame = false
		}

	}

	if allSame {
		return 0
	}

	cnt1Total := len(S) - cnt0Total

	cnt1 := 0
	cnt0 := 0
	for _, v := range S {
		flipCnt := cnt1 + cnt0Total - cnt0

		if minFlip == -1 || flipCnt < minFlip {
			minFlip = flipCnt
		}
		if v == '1' {

			cnt1++
		}
		if v == '0' {
			cnt0++
		}

	}

	if cnt0Total < minFlip {
		return cnt0Total
	}

	if cnt1Total < minFlip {
		return cnt1Total
	}

	return minFlip

}
