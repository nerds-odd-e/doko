package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FizzBuzz_input_1_output_should_be_1(t *testing.T) {
	result := fizzBuzz(1)
	assert.Equal(t, result, "1", "input is 1, so we should get 1")
}
