package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FizzBuzz_input_1(t *testing.T) {
	result := fizzBuzz(1)
	assert.Equal(t, result, "1", "input is 1, so we should get 1")
}

func Test_FizzBuzz_input_3(t *testing.T) {
	result := fizzBuzz(3)
	assert.Equal(t, result, "Fizz", "Input is 3 which is Fizz, so we should get Fizz")
}

func Test_FizzBuzz_input_5(t *testing.T) {
	result := fizzBuzz(5)
	assert.Equal(t, result, "Buzz", "Input is 5 which is Buzz, so we should get Buzz")
}
