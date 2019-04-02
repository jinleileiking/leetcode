package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	assert.Equal(t, []int{}, findSubstring("a", []string{}))
	assert.Equal(t, []int{8}, findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
	assert.Equal(t, []int{0, 9}, findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	assert.Equal(t, []int{}, findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))

}
