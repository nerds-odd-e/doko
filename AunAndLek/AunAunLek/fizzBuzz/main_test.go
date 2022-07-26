package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {

	type testcase struct {
		expected string
		input    int
	}
	cases := []testcase{
		{expected: "1", input: 1},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("Input %v get '%v'", c.input, c.expected), func(t *testing.T) {
			assert.Equal(t, c.expected, FizzBuzz(c.input))
		})
	}
	expected := cases[0].expected
	input := cases[0].input
	t.Run(fmt.Sprintf("Input %v get '%v'", input, expected), func(t *testing.T) {
		assert.Equal(t, expected, FizzBuzz(input))
	})

	t.Run("Input 2 get '2'", func(t *testing.T) {
		assert.Equal(t, "2", FizzBuzz(2))
	})

	t.Run("Input 3 get 'fizz'", func(t *testing.T) {
		assert.Equal(t, "fizz", FizzBuzz(3))
	})

	t.Run("Input 5 get 'buzz'", func(t *testing.T) {
		assert.Equal(t, "buzz", FizzBuzz(5))
	})

	t.Run("Input 6 get 'fizz'", func(t *testing.T) {
		assert.Equal(t, "fizz", FizzBuzz(6))
	})

	t.Run("Input 10 get 'buzz'", func(t *testing.T) {
		assert.Equal(t, "buzz", FizzBuzz(10))
	})
}
