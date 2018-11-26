package main

func isNStraightHand(hand []int, W int) bool {

	if W == 1 {
		return true
	}
	// sort the hand

	sorted := []int{}
	var ret bool

	sorted = Heapsort(hand)

	// spew.Dump(sorted)
	for {

		//get W

		ret, sorted = getW(sorted, W)
		// spew.Dump(ret, sorted)

		if !ret {
			return false
		}

		if len(sorted) == 0 {
			break
		}
	}

	return true

}

func getW(arr []int, W int) (bool, []int) {

	var last int
	var gotCnt int
	ret := []int{}
	last = arr[0]
	arr = arr[1:]
	for k, v := range arr {
		// spew.Dump(k, v, ret)
		if last == v {
			ret = append(ret, v)
			continue
		}
		if v != last-1 {
			return false, nil
		}
		//got it
		last = v
		gotCnt++
		if gotCnt == W-1 {
			// spew.Dump(arr[k+1:])
			// add tail
			if k < len(arr)-1 {
				ret = append(ret, arr[k+1:]...)
			}
			return true, ret
		}

	}
	return false, nil
}

func Heapsort(arr []int) []int {
	final := []int{}
	for len(arr) != 0 {
		ret := Heapify(arr)
		final = append(final, ret[0])
		arr = arr[1:]
	}
	return final
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
