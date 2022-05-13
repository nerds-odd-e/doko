package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivisbleByThreeReturnFizz(t *testing.T) {
	assert.Equal(t, "Fizz", FizzBuzz(3))
	assert.Equal(t, "Fizz", FizzBuzz(6))
}

func TestFourReturnFour(t *testing.T) {
	assert.Equal(t, "4", FizzBuzz(4))
}

func TestFiveReturnBuzz(t *testing.T) {
	assert.Equal(t, "Buzz", FizzBuzz(5))
}

func TestSevenReturnSeven(t *testing.T) {
	assert.Equal(t, "7", FizzBuzz(7))
}
