package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalNumber(t *testing.T) {
	assert.Equal(t, "1", findFizzBuzz(1))
}

func TestFizz(t *testing.T) {
	assert.Equal(t, "Fizz", findFizzBuzz(3))
}

func TestBuzz(t *testing.T) {
	assert.Equal(t, "Buzz", findFizzBuzz(5))
}
