package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	assert.Equal(t, 1, minIncrementForUnique([]int{1, 2, 2}))
	assert.Equal(t, 6, minIncrementForUnique([]int{3, 2, 1, 2, 1, 7}))
}

func TestMap(t *testing.T) {

	// 1, -> 1, 2,  ---> 1,2,3  == 2
	m := make(map[int]bool)

	m[1] = true
	m[2] = true

	assert.Equal(t, 5, checkMap(1, m, 3))
	spew.Dump(m)
}
