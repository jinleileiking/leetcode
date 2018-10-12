package main

import (
	"strconv"
)

func decodeAtIndex(S string, K int) string {

	var temp string
	temp = S
	for {

		str, toFindInt := findStr(temp)
		i, left := findInt(toFindInt)
		if str == "" || i == nil {
			break
		}
		temp = genStr(str, *i, left)

		if len(temp) >= K {
			return temp[K-1 : K]
		}
	}

	return ""
}

func findStr(str string) (string, string) {
	var ret string
	var lastpos int
	var intFound bool
	for p, c := range str {
		lastpos = p
		if c >= 'A' && c <= 'z' {
			ret = ret + string(c)
		} else {
			intFound = true
			break
		}
	}
	if lastpos == len(str)-1 && !intFound {
		return ret, ""
	}

	return ret, str[lastpos:]
}

func findInt(str string) (*int, string) {
	var ret string
	var lastpos int
	for p, c := range str {
		lastpos = p
		if c >= '0' && c <= '9' {
			ret = ret + string(c)
		} else {
			break
		}
	}

	i, err := strconv.Atoi(ret)

	if err != nil {
		return nil, str[lastpos:]
	}

	if lastpos == len(str)-1 {
		return &i, ""
	}

	return &i, str[lastpos:]
}

func genStr(str string, i int, left string) string {
	var ret string
	for t := 0; t < i; t++ {
		ret = ret + str
	}
	return ret + left
}
