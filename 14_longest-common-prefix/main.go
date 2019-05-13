package main

import "sync"

var wg sync.WaitGroup

func longestCommonPrefix(strs []string) string {

	// corner case

	if len(strs) == 1 {
		return strs[0]
	}

	for {
		parsed := parse(strs)

		if len(parsed) == 1 {
			return parsed[0]
		}
	}

}

func parse(strs []string) []string {

	parsed := []string{}
	for p := 0; p <= len(strs)-1; p += 2 {
		wg.Add(1)
		go func() {
			parsed = append(parsed, parsePre(strs[p:p+1]))
			wg.Done()
		}()
	}

	if len(strs)%2 != 0 {
		parsed = append(parsed, strs[len(strs)-1])
	}

	wg.Wait()

	return parsed

}

func parsePre(str2 []string) string {
	var ret string
	for k, v := range str2[0] {
		if k <= len(str2)-1 && byte(v) == str2[1][k] {
			ret = ret + string(v)
		} else {
			break
		}
	}

	return ret
}
