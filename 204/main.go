package main

func countPrimes(n int) int {

	if n == 0 || n == 1 || n == 2 {
		return 0
	}
	if n == 0 || n == 1 || n == 3 {
		return 1
	}
	var ret []int

	for i := 0; i < n; i++ {
		ret = append(ret, i)
	}

	// spew.Dump(ret)

	for i := 2; i < n; i++ {
		for j := 0; j < len(ret); j++ {
			// spew.Dump(ret[j], j, "ssssssssssssssssssss")
			if ret[j] == 1 || ret[j] == 0 {
				ret = append(ret[:j], ret[j+1:]...)
			}
			if ret[j]%i == 0 && i != ret[j] {
				// spew.Dump("cccccccccccccccc", i, ret[j])
				ret = append(ret[:j], ret[j+1:]...)
			}
		}
	}

	// spew.Dump(ret)
	return len(ret)
}
