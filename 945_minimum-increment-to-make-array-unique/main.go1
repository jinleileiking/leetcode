package main

func minIncrementForUnique(A []int) int {
	cnt := 0

	changed := true
	for {
		changed, A = findAndChange(A)
		if changed {
			cnt++
		} else {
			break
		}

	}

	return cnt
}

func findAndChange(A []int) (bool, []int) {
	for k1, v1 := range A {
		for k2 := k1 + 1; k2 < len(A); k2++ {
			if v1 == A[k2] {
				A[k2] = A[k2] + 1
				return true, A
			}
		}
	}

	return false, A
}
