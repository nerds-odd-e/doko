package fizzbuzz_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"tdd.com/v1/pong-chet/fizzbuzz"
)

func Test_FizzBuzz_say_the_number_by_default(t *testing.T) {
	assert.Equal(t, "1", fizzbuzz.FizzBuzz(1))
	assert.Equal(t, "2", fizzbuzz.FizzBuzz(2))
}

func Test_FizzBuzz_say_Fizz_when_divisible_by_3(t *testing.T) {
	assert.Equal(t, "Fizz", fizzbuzz.FizzBuzz(3))
	assert.Equal(t, "Fizz", fizzbuzz.FizzBuzz(6))
}

func Test_FizzBuzz_6_Fizz(t *testing.T) {
}

func Test_FizzBuzz_5_Buzz(t *testing.T) {
	assert.Equal(t, "Buzz", fizzbuzz.FizzBuzz(5))
}

func Test_FizzBuzz_10_Buzz(t *testing.T) {
	assert.Equal(t, "Buzz", fizzbuzz.FizzBuzz(10))
}

func Test_FizzBuzz_15_FizzBuzz(t *testing.T) {
	assert.Equal(t, "FizzBuzz", fizzbuzz.FizzBuzz(15))
}

func Test_FizzBuzz_30_FizzBuzz(t *testing.T) {
	assert.Equal(t, "FizzBuzz", fizzbuzz.FizzBuzz(30))
}
