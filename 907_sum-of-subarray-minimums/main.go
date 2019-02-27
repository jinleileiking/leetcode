package main

func mins(arr []int) (int, []int) {
	var min int
	minK := []int{}
	min = arr[0]

	for k, v := range arr {
		if v == 0 {
			continue
		}
		if v < min || (v > min && min == 0) {
			min = v
			minK = []int{}
			minK = append(minK, k)
		} else if v == min {
			min = v
			minK = append(minK, k)
		}
	}

	return min, minK
}

func findEdge(A []int, pos int) (left, right int) {

	left = pos
	right = pos
	var p int

	got0 := false
	for p = pos; p >= 0; p-- {
		if A[p] == 0 {
			// left = p
			got0 = true
			break
		}
		left = p
	}

	if p == 0 && !got0 {
		left = 0
	}

	for p = pos; p < len(A); p++ {
		if A[p] == 0 {
			// right = p
			break
		}
		right = p
	}
	if p == len(A)-1 {
		right = len(A) - 1
	}

	return
}

func calc(pos, left, right, m int) int {
	return m * ((pos - left) * (right - pos + 1))
}

func sumSubarrayMins(A []int) int {

	ret := 0
	for {
		m, mPoses := mins(A)

		if m == 0 {
			break
		}

		for _, pos := range mPoses {
			left, right := findEdge(A, pos)
			ret += calc(pos, left, right, m)
		}

		for _, pos := range mPoses {
			A[pos] = 0
		}
	}

	return ret
}
