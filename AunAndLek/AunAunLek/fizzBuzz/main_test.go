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
		{expected: "2", input: 2},
		{expected: "fizz", input: 3},
		{expected: "buzz", input: 5},
		{expected: "fizz", input: 6},
		{expected: "buzz", input: 10},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("Input %v get '%v'", c.input, c.expected), func(t *testing.T) {
			assert.Equal(t, c.expected, FizzBuzz(c.input))
		})
	}
}
