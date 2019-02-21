package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	// assert.Equal(t, "1230", multiply("123", "10"))
	// assert.Equal(t, "144", multiply("12", "12"))
	assert.Equal(t, "56088", multiply("123", "456"))
}

func TestReverse(t *testing.T) {

	assert.Equal(t, "7654321", reverse("1234567"))
}

func TestSplitInt(t *testing.T) {

	assert.Equal(t, []int{5, 4, 3, 2, 1}, splitInt(12345))
}
