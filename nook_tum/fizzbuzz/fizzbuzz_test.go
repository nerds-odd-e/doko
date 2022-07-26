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

func TestInput5_ShouldSayBuzz(t *testing.T) {
	answer := FizzBuzz(5)
	expected := "Buzz"
	assert.Equal(t, answer, expected, "Wanted %s but got %s", expected, answer)
}

func TestInput15_ShouldSayFizzBuzz(t *testing.T) {
	answer := FizzBuzz(15)
	expected := "1"
	assert.Equal(t, answer, expected, "Wanted FizzBuzz but got %s", answer)
}
