package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	assert.Equal(t, "aaa", longestCommonPrefix([]string{"aaab", "aaac"}))

}
