package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	assert.Equal(t, 1, minFlipsMonoIncr("00110"))
	assert.Equal(t, 2, minFlipsMonoIncr("010110"))
	assert.Equal(t, 2, minFlipsMonoIncr("00011000"))
	assert.Equal(t, 0, minFlipsMonoIncr("11111"))
	assert.Equal(t, 0, minFlipsMonoIncr("011111"))
	assert.Equal(t, 3, minFlipsMonoIncr("1111001110"))

}
