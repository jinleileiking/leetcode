package main

func findSubstring(s string, words []string) []int {

	//corner case

	if len(words) == 0 {
		return []int{}
	}

	ret := []int{}
	for p, _ := range s {
		if true == check(s[p:], words) {
			ret = append(ret, p)
		}
	}

	return ret

}

func check(s string, words []string) bool {
	m := make(map[string]int)

	for _, v := range words {
		m[v]++
		// if _, ok := m[v]; ok{
		// 	m[v]++
		// }else{
		// 	m[v]++
		// }
	}

	loop := 0
	lenWord := len(words[0])
	for i := 0; i <= len(s)-lenWord; i = i + lenWord {
		item := s[i : i+lenWord]
		if cnt, ok := m[item]; ok {
			m[item] = cnt - 1

			if m[item] < 0 {
				return false
			}
			loop++
			if loop == len(words) {
				break
			}
		} else {
			return false
		}
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true

}
