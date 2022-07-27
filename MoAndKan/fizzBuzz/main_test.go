package fizzBuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzzWithInputThree(t *testing.T) {
	result := fizzBuzz(3)
	assert.Equal(t, result, "fizz", "Input is 3, should get fizz")
}

func TestFizzBuzzWithInputFive(t *testing.T) {
	result := fizzBuzz(5)
	assert.Equal(t, result, "Buzz", "Input is 5, should get Buzz")
}

type TestCase struct {
	name       string
	input      int
	expected   string
	errMessage string
}

func TestFizz(t *testing.T) {
	testCases := []TestCase{
		TestCase{"Input is 1", 1, "1", "Input is 1, should get 1"},
	}

	for _, testCase := range testCases {
		result := fizzBuzz(testCase.input)
		assert.Equal(t, result, testCase.expected, testCase.errMessage)
	}
}
