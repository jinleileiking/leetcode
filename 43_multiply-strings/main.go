package main

import (
	"strconv"
)

func reverse(src string) string {
	var dst string
	for _, v := range src {
		dst = string(v) + dst
	}

	return dst
}
func multiply(num1 string, num2 string) string {

	if num1 == "0" || num2 == "0" {
		return "0"
	}
	num1 = reverse(num1)
	num2 = reverse(num2)

	dst := make(map[int]int)
	for k1, v1 := range num1 {
		for k2, v2 := range num2 {
			// spew.Dump(k1, k2)
			v1Int, _ := strconv.Atoi(string(v1))

			v2Int, _ := strconv.Atoi(string(v2))

			mul := v1Int * v2Int

			high := mul / 10
			low := mul % 10

			// if low != 0 {
			dst[k1+k2] = dst[k1+k2] + low
			// }
			if high != 0 {
				dst[k1+k2+1] = dst[k1+k2+1] + high
			}

		}
	}
	// spew.Dump(dst)

	needRepeat := true
	for needRepeat {
		needRepeat = false
		// spew.Dump(dst)
		for pos, val := range dst {
			if val >= 10 {
				needRepeat = true
				splited := splitInt(val)
				// spew.Dump(splited)
				for k, v := range splited {
					if k == 0 {
						dst[pos] = v
						continue
					}
					dst[pos+k] = dst[pos+k] + v
				}

				break
			}
		}
		// spew.Dump(dst)
		// os.Exit(0)
	}

	// spew.Dump(dst)

	// ret := []byte{}

	lens := getMaxKey(dst)
	ret := make([]byte, lens+1)
	for pos, val := range dst {
		// if pos == 0 {
		// 	continue
		// }
		// spew.Dump(string(strconv.Itoa(val)[0]))
		ret[pos] = byte(strconv.Itoa(val)[0])
	}

	// if string(ret) == "\x00" || string(ret) == "" {
	// 	return "0"
	// }
	return reverse(string(ret))
}

func getMaxKey(m map[int]int) int {
	var max int
	for k, _ := range m {
		if max < k {
			max = k
		}
	}

	return max
}

func splitInt(i int) []int {
	ret := []int{}
	for i != 0 {
		ret = append(ret, i%10)
		i = i / 10
	}

	return ret
}
