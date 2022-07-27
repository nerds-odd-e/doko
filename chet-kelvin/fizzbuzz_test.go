package chet_kelvin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInput1(t *testing.T) {
	assert.Equal(t, "1", fizzbuzz(1))
}

func TestInput2(t *testing.T) {
	assert.Equal(t, "2", fizzbuzz(2))
}

func TestInput3(t *testing.T) {
	assert.Equal(t, "Fizz", fizzbuzz(3))
}

func TestInput6(t *testing.T) {
	assert.Equal(t, "Fizz", fizzbuzz(6))
}
