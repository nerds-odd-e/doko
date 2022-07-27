package fizzbuzz_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"tdd.com/v1/pong-chet/fizzbuzz"
)

func Test_FizzBuzz_input_1_output_1(t *testing.T) {

	result := fizzbuzz.FizzBuzz(1)

	assert.Equal(t, "1", result)

}
