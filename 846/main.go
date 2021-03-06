package main

func isNStraightHand(hand []int, W int) bool {

	if W == 1 {
		return true
	}

	tmpArr := hand

	for {
		// spew.Dump(tmpArr)
		// get max
		max := GetMax(tmpArr)
		// spew.Dump(max)

		if hasW(hand, W, max) {
			tmpArr = cutW(tmpArr, W, max)
		} else {
			return false
		}

		if len(tmpArr) == 0 {
			return true
		}

		if len(tmpArr) < W {
			return false
		}
	}

}
func hasW(arr []int, W int, max int) bool {

	// ret := []int{}

	for i := 0; i < W; i++ {
		var found bool
		for _, v := range arr {
			if v == max-i {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

func cutW(arr []int, W int, max int) []int {

	ret := []int{}
	got := []int{}

	for _, v := range arr {
		// spew.Dump(v, ret)
		if v <= max && v > max-W {
			goted := false
			for _, gotItem := range got {
				if v == gotItem {
					goted = true
					break
				}
			}

			if goted {
				ret = append(ret, v)
				continue
			}
			got = append(got, v)
			continue
		}

		ret = append(ret, v)
	}

	return ret
}

func GetMax(arr []int) int {
	return Heapify(arr)[0]
}

func Heapify(arr []int) []int {

	for i := len(arr) - 1; i > 0; {
		// spew.Dump(arr)
		// spew.Dump("aaaaaaaaaa")
		if change(&arr, i) {
			i = len(arr) - 1
			// spew.Dump(arr)
			// continue
		} else {
			if i%2 == 0 {
				i = i - 2
			} else {
				i = i - 1
			}
		}
	}

	return arr
}

func change(arr *[]int, pos int) bool {

	in := *arr
	parentPos := (pos - 1) / 2
	leftPos := pos - 1
	rightPos := pos

	right := in[rightPos]
	left := in[leftPos]
	parent := in[parentPos]

	// one leaf
	if pos%2 != 0 {
		// just check self and his parent

		if parent < in[pos] {
			self := in[pos]
			in[pos] = parent
			in[parentPos] = self
			*arr = in
			return true
		}

	}

	if parent < left && left >= right {
		in[parentPos] = left
		in[leftPos] = parent
		*arr = in
		return true
	} else if parent < right && right >= left {
		in[parentPos] = right
		in[rightPos] = parent
		*arr = in
		return true
	}

	*arr = in
	return false
}
