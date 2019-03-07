package main

func minIncrementForUnique(A []int) int {

	m := make(map[int]bool)
	cnt := 0
	for _, v := range A {
		cnt = checkMap(v, m, cnt)
	}

	return cnt
}

func checkMap(a int, m map[int]bool, cnt int) int {
	if found := m[a]; found {
		cnt++
		return checkMap(a+1, m, cnt)
	} else {
		m[a] = true
	}

	return cnt
}
