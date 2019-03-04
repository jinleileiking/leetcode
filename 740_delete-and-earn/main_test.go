package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	assert.Equal(t, 6, deleteAndEarn([]int{3, 4, 2}))
	assert.Equal(t, 9, deleteAndEarn([]int{2, 2, 3, 3, 3, 4}))
}
