package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalNumber(t *testing.T) {
	assert.Equal(t, "1", findFizzBuzz(1))
	assert.Equal(t, "2", findFizzBuzz(2))
	assert.Equal(t, "4", findFizzBuzz(4))
	assert.Equal(t, "6", findFizzBuzz(6))
}

func TestFizz(t *testing.T) {
	assert.Equal(t, "Fizz", findFizzBuzz(3))
}

func TestBuzz(t *testing.T) {
	assert.Equal(t, "Buzz", findFizzBuzz(5))
}

func TestFizzBuzz(t *testing.T) {
	assert.Equal(t, "FizzBuzz", findFizzBuzz(15))
}
