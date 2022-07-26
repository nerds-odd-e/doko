package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FizzBuzz_input_1_output_should_be_1(t *testing.T) {
	result := fizzBuzz(1)
	assert.Equal(t, result, "1", "input is 1, so we should get 1")
}

func Test_FizzBuzz_input_3_output_should_be_Fizz(t *testing.T) {
	result := fizzBuzz(3)
	assert.Equal(t, result, "Fizz", "Input is 3 which is Fizz, so we should get Fizz")
}

func Test_FizzBuzz_input_5_output_should_be_Buzz(t *testing.T) {
	result := fizzBuzz(5)
	assert.Equal(t, result, "5", "Input is 5 which is Buzz, so we should get Buzz")
}
