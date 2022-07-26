package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInput1_ShouldSay1(t *testing.T) {
	answer := FizzBuzz(1)
	expected := "1"
	assert.Equal(t, answer, expected, "")
}

func TestInput1_ShouldSay3(t *testing.T) {
	answer := FizzBuzz(3)
	expected := "Fizz"
	assert.Equal(t, answer, expected, "Fizz")
}
