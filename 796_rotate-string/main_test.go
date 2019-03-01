package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	assert.Equal(t, true, rotateString("abcde", "cdeab"))
	assert.Equal(t, false, rotateString("abcde", "abced"))
}
