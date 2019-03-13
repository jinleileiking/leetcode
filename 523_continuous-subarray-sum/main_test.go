package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	assert.Equal(t, true, checkSubarraySum([]int{23, 2, 4, 6, 7}, 6))
	assert.Equal(t, true, checkSubarraySum([]int{23, 2, 4, 6, 7}, 42))
	assert.Equal(t, false, checkSubarraySum([]int{5, 2, 4}, 5))
}
