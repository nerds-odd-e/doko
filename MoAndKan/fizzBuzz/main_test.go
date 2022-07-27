package fizzBuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	name       string
	input      int
	expected   string
	errMessage string
}

func TestFizzBuzz(t *testing.T) {
	testCases := []TestCase{
		TestCase{"Input is 1", 1, "1", "Input is 1, should get 1"},
		TestCase{"Input is 3", 3, "fizz", "Input is 3, should get fizz"},
		TestCase{"Input is 5", 5, "Buzz", "Input is 5, should get Buzz"},
		TestCase{"Input is 6", 6, "6", "Input is 6, should get fizz"},
	}

	for _, testCase := range testCases {
		result := fizzBuzz(testCase.input)
		assert.Equal(t, result, testCase.expected, testCase.errMessage)
	}
}
