package fizzbuzz_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"tdd.com/v1/lek-pong/fizzbuzz"
)

type TestCase struct {
	name     string
	input    int
	expected string
}

func TestFizzBuzz(t *testing.T) {
	cases := []TestCase{
		{name: "fizzbuzz 1", input: 1, expected: "1"},
		{name: "fizzbuzz 2", input: 2, expected: "2"},
		{name: "fizzbuzz 3", input: 3, expected: "3"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, fizzbuzz.FizzBuzz(c.input))
		})
	}
}
