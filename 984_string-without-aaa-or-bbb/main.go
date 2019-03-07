package main

func strWithout3a3b(A int, B int) string {

	var ret string

	for {
		if A == 0 {
			if B == 1 {
				ret = ret + "b"
			}
			if B == 2 {
				ret = ret + "bb"
			}
			return ret
		}
		if B == 0 {
			if A == 1 {
				ret = ret + "a"
			}
			if A == 2 {
				ret = ret + "aa"
			}
			return ret
		}

		if A-B > 1 {
			ret = ret + "aa"
			ret = ret + "b"
			A = A - 2
			B = B - 1
			continue
		} else if B-A > 1 {
			ret = ret + "bb"
			ret = ret + "a"
			B = B - 2
			A = A - 1
			continue
		} else {
			if B > A {
				for i := 0; i < (B-1)/2; i++ {
					ret = ret + "bb"
					ret = ret + "aa"
				}

				if (B-1)%2 != 0 {
					ret = ret + "b"
					ret = ret + "a"
				}

			} else if B < A {
				for i := 0; i < (A-1)/2; i++ {
					ret = ret + "aa"
					ret = ret + "bb"
				}
				if (A-1)%2 != 0 {
					ret = ret + "a"
					ret = ret + "b"
				}
			} else {
				for i := 0; i < B/2; i++ {
					ret = ret + "bb"
					ret = ret + "aa"
				}

				if B%2 != 0 {
					ret = ret + "b"
					ret = ret + "a"
				}
			}
			if B > A {
				if (B-A)%2 != 0 {
					ret = ret + "b"
				}
			}
			if A > B {
				if (A-B)%2 != 0 {
					ret = ret + "a"
				}
			}
			return ret
		}
	}
}

// var isALong bool
// var ret, bigStr, smallStr string
// var big, small int

// diff := 0

// if A > B {
// 	diff = A - B
// 	big = A
// 	small = B
// 	bigStr = "a"
// 	smallStr = "b"

// } else {
// 	diff = B - A
// 	isALong = false
// 	big = B
// 	small = A
// 	bigStr = "b"
// 	smallStr = "a"
// }

// for i := 0; i < (big-diff)/2; i++ {
// 	ret = ret + bigStr + bigStr
// 	ret = ret + smallStr + smallStr
// }

// if (big-diff)%2 != 0 {
// 	ret = ret + bigStr
// 	ret = ret + smallStr + smallStr
// }
