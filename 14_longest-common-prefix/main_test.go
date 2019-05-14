package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	// assert.Equal(t, "aaa", longestCommonPrefix([]string{"aaab", "aaac"}))
	// assert.Equal(t, "aa", longestCommonPrefix([]string{"aa", "aaa"}))
	// assert.Equal(t, "aa", longestCommonPrefix([]string{"aaab", "aaac", "aaaa", "aabb"}))
	// assert.Equal(t, "fl", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	assert.Equal(t, "", longestCommonPrefix([]string{}))

}

func TestParse(t *testing.T) {

	assert.Equal(t, []string{"aa"}, parse([]string{"aaa", "aa"}))

}
