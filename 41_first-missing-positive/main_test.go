package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	assert.Equal(t, 4, firstMissingPositive([]int{1, 2, 3}))
	assert.Equal(t, 2, firstMissingPositive([]int{3, 1, -1, 4}))
	assert.Equal(t, 1, firstMissingPositive([]int{2147483647}))
}
