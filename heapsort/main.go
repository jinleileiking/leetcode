package main

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
