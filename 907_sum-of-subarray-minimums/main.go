package main

func minAndModify(arr []int) (int, int, []int) {
	var min, minK int
	min = arr[0]

	for k, v := range arr {
		if v == 0 {
			continue
		}
		if v < min {
			min = v
			minK = k
		}
	}

	t := append(arr[:minK], 0)
	return min, minK, append(t, arr[minK+1:]...)
}

// func findSubArr() {
// }

func sumSubarrayMins(A []int) int {

	src := A

	data := make(map[int]int)
	for _, v := range A {
		m, mPos, src := minAndModify(src)

		data[m] = data[m] + 1
		// src = append(src[

		if len(src) == 0 {
			break
		}
	}

	return 0
}
