package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FizzBuzz_input_1_output_should_be_1(t *testing.T) {
	result := fizzBuzz(1)
	assert.Equal(t, result, "1", "input 1 should be 1")
}
