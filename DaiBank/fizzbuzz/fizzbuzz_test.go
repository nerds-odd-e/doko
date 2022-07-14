package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalNumber1(t *testing.T) {
	assert.Equal(t, "1", findFizzBuzz(1))
}

func TestNormalNumber2(t *testing.T) {
	assert.Equal(t, "2", findFizzBuzz(2))
}

func TestNormalNumber3(t *testing.T) {
	assert.Equal(t, "Fizz", findFizzBuzz(3))
}
func TestNormalNumber4(t *testing.T) {
	assert.Equal(t, "4", findFizzBuzz(4))
}
func TestNormalNumber5(t *testing.T) {
	assert.Equal(t, "Buzz", findFizzBuzz(5))
}
func TestNormalNumber6(t *testing.T) {
	assert.Equal(t, "Fizz", findFizzBuzz(6))
}

func TestNormalNumber15(t *testing.T) {
	assert.Equal(t, "FizzBuzz", findFizzBuzz(15))
}

func TestFizzBuzz(t *testing.T) {
	assert.Equal(t, "Fizz", findFizzBuzz(30))
}
