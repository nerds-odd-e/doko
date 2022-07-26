package fizzbuzz_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"tdd.com/v1/pong-chet/fizzbuzz"
)

func Test_FizzBuzz_1_1(t *testing.T) {
	assert.Equal(t, "1", fizzbuzz.FizzBuzz(1))
}

func Test_FizzBuzz_2_2(t *testing.T) {
	assert.Equal(t, "2", fizzbuzz.FizzBuzz(2))
}

func Test_FizzBuzz_3_Fizz(t *testing.T) {
	assert.Equal(t, "3", fizzbuzz.FizzBuzz(3))
}
