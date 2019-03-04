package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	assert.Equal(t, 1, deleteAndEarn([]int{1, 2, 2}))
	assert.Equal(t, 6, deleteAndEarn([]int{3, 2, 1, 2, 1, 7}))
}
