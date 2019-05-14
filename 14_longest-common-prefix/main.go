package main

import (
	"sync"
)

var wg sync.WaitGroup

func longestCommonPrefix(strs []string) string {

	// corner case
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	parsed := strs
	for {
		parsed = parse(parsed)

		if len(parsed) == 1 {
			return parsed[0]
		}
	}

}

func parse(strs []string) []string {

	parsed := []string{}
	for p := 0; p <= len(strs)-1-len(strs)%2; p += 2 {
		str2 := strs[p : p+2]
		wg.Add(1)
		go func([]string) {
			parsed = append(parsed, parsePre(str2))
			wg.Done()
		}(str2)
	}

	if len(strs)%2 != 0 {
		parsed = append(parsed, strs[len(strs)-1])
	}

	wg.Wait()

	return parsed

}

func parsePre(str2 []string) string {
	if len(str2) > 2 {
		panic(0)
	}

	var ret string
	str0 := str2[0]
	str1 := str2[1]

	for k, v := range str0 {
		if k <= len(str1)-1 && byte(v) == str1[k] {
			ret = ret + string(v)
		} else {
			break
		}
	}

	return ret
}
