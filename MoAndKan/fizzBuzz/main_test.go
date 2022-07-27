package fizzBuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	result := fizzBuzz(1)
	assert.Equal(t, result, "1", "Input is 1, should get 1")
}
