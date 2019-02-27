package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	assert.Equal(t, 17, sumSubarrayMins([]int{3, 1, 2, 4}))
}

func TestFunc(t *testing.T) {

	// m, p, a := minAndModify([]int{7, 6, 3, 0, 4, 5})
	// assert.Equal(t, 3, m) // assert.Equal(t, 2, p)
	// assert.Equal(t, []int{7, 6, 0, 0, 4, 5}, a)
}

func TestMins(t *testing.T) {

	var m int
	var poses []int
	// m, poses = mins([]int{5, 2, 3, 2, 4})
	// assert.Equal(t, 2, m)
	// assert.Equal(t, []int{1, 3}, poses)

	// m, poses = mins([]int{5, 2, 0, 2, 4})
	// assert.Equal(t, 2, m)
	// assert.Equal(t, []int{1, 3}, poses)

	// m, poses = mins([]int{0, 2, 0, 2, 0})
	// assert.Equal(t, 2, m)
	// assert.Equal(t, []int{1, 3}, poses)

	// m, poses = mins([]int{2})
	// assert.Equal(t, 2, m)
	// assert.Equal(t, []int{0}, poses)

	// m, poses = mins([]int{2, 0})
	// assert.Equal(t, 2, m)
	// assert.Equal(t, []int{0}, poses)

	m, poses = mins([]int{3, 1, 2, 4})
	assert.Equal(t, 1, m)
	assert.Equal(t, []int{1}, poses)

}

func TestEdge(t *testing.T) {
	var left, right int
	// left, right = findEdge([]int{5, 2, 3, 2, 4}, 1)
	// assert.Equal(t, 0, left)
	// assert.Equal(t, 4, right)

	// left, right = findEdge([]int{0, 5, 2, 3, 2, 0, 4}, 1)
	// assert.Equal(t, 1, left)
	// assert.Equal(t, 4, right)

	// left, right = findEdge([]int{0}, 0)
	// assert.Equal(t, 0, left)
	// assert.Equal(t, 0, right)

	left, right = findEdge([]int{3, 1, 2, 4}, 1)
	assert.Equal(t, 0, left)
	assert.Equal(t, 3, right)
}

func TestCalc(t *testing.T) {
	assert.Equal(t, 6, calc(1, 0, 3, 1))
}
