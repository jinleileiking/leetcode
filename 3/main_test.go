package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestLong(t *testing.T) {

	spew.Dump(lengthOfLongestSubstring("abcabcbb"))
}

func TestMain(t *testing.T) {

	spew.Dump(find("abcabcbb", 0))
}

func TestLong1(t *testing.T) {

	assert.Equal(t, 3, lengthOfLongestSubstring("abc"), "they should be equal")
}
func TestLong2(t *testing.T) {

	assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"), "they should be equal")

}

func TestMain2(t *testing.T) {

	assert.Equal(t, 1, find("bbbbb", 0), "bbbbb they should be equal")
	assert.Equal(t, 1, find("bbbb", 1), "bbbb they should be equal")
	assert.Equal(t, 1, find("bbb", 1), "bbb they should be equal")
	assert.Equal(t, 1, find("bb", 1), "bb they should be equal")
	assert.Equal(t, 1, find("b", 1), "b they should be equal")
}

func TestLong3(t *testing.T) {

	assert.Equal(t, 2, lengthOfLongestSubstring("aab"), "they should be equal")
}
