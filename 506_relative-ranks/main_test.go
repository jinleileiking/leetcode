package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	assert.Equal(t, []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"},
		findRelativeRanks([]int{5, 4, 3, 2, 1}))
}
