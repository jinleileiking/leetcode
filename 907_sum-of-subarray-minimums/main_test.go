package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	// assert.Equal(t, "0", multiply("123", "0"))
}

func TestFunc(t *testing.T) {

	m, p, a := minAndModify([]int{7, 6, 3, 0, 4, 5})
	assert.Equal(t, 3, m)
	assert.Equal(t, 2, p)
	assert.Equal(t, []int{7, 6, 0, 0, 4, 5}, a)
}
