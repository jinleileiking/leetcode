package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	assert.Equal(t, [][]int{{-1, 0, 1}, {-1, -1, 2}}, threeSum([]int{-1, 0, 1, 2, -1, -4}))
	assert.Equal(t, [][]int{}, threeSum([]int{1, 2, -2, -1}))

}
