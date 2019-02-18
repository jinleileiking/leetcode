package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	assert.Equal(t, true, isMatch("aaa", "aaa"))
	assert.Equal(t, false, isMatch("aaa", "bb"))
	assert.Equal(t, true, isMatch("aaa", "a.."))
	assert.Equal(t, false, isMatch("aaaaa", "a..a"))
	assert.Equal(t, true, isMatch("aaaaa", "a..a."))
	assert.Equal(t, true, isMatch("aaaaa", "a..aa"))
	assert.Equal(t, true, isMatch("aaaaa", "a*aa"))
	assert.Equal(t, true, isMatch("aaaaa", "a*a."))
	assert.Equal(t, false, isMatch("abcaa", "a*a."))
	assert.Equal(t, true, isMatch("aaa", "a*a."))
	assert.Equal(t, true, isMatch("aaaaa", "a*a."))
	assert.Equal(t, false, isMatch("aaaaa", "*a."))
	assert.Equal(t, true, isMatch("aaaaa", "a*a."))
	assert.Equal(t, true, isMatch("abcde", ".*"))
	assert.Equal(t, true, isMatch("aab", "a*b"))
	assert.Equal(t, true, isMatch("aab", "c*a*b"))
	assert.Equal(t, true, isMatch("ippi", "ip*."))
	assert.Equal(t, true, isMatch("iss", "is*"))
	assert.Equal(t, true, isMatch("ss", "s*"))
	assert.Equal(t, true, isMatchInner("ssi", "s*i"))
	assert.Equal(t, true, isMatchInner("ssiss", "s*is*"))
	assert.Equal(t, true, isMatchInner("ssis", "s*is"))
	assert.Equal(t, true, isMatch("mississippi", "mis*is*ip*."))
	assert.Equal(t, false, isMatch("ab", ".*c"))
	assert.Equal(t, false, isMatch("aaaa", "aaaaa"))
}
