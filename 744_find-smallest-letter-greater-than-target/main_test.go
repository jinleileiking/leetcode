package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	assert.Equal(t, byte('c'), nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a'))
}
