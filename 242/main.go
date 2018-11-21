package main

import "github.com/davecgh/go-spew/spew"

func isAnagram(s string, t string) bool {

	var sstat, dstat map[int]int

	sstat = make(map[int]int)
	dstat = make(map[int]int)

	for _, c := range s {
		sstat[int(c)]++
	}

	for _, c := range t {
		dstat[int(c)]++
	}

	spew.Dump(sstat, dstat)

	if len(sstat) > len(dstat) {
		for k, count := range sstat {
			if dstat[k] != count {
				return false
			}
		}
	} else {
		for k, count := range dstat {
			if sstat[k] != count {
				return false
			}
		}
	}

	return true
}
