package fizzbuzz_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"tdd.com/v1/jade_hok/fizzbuzz"
)

func Test_FizzBuzzInput1Output1(t *testing.T) {
	assert.Equal(t, "1", fizzbuzz.FizzBuzz(1))
}

func TestFizzBuzzInput2Output2(t *testing.T) {
	assert.Equal(t, "2", fizzbuzz.FizzBuzz(2))
}

func TestFizzBuzzInput3OutputFizz(t *testing.T) {
	assert.Equal(t, "fizz", fizzbuzz.FizzBuzz(3))
}

func TestFizzBuzzInput5OutputBuzz(t *testing.T) {
	assert.Equal(t, "buzz", fizzbuzz.FizzBuzz(5))
}

func TestFizzBuzzInput6OutputFizz(t *testing.T) {
	assert.Equal(t, "fizz", fizzbuzz.FizzBuzz(6))
}

func TestFizzBuzzInput10OutputBuzz(t *testing.T) {
	assert.Equal(t, "buzz", fizzbuzz.FizzBuzz(10))
}

func TestFizzBuzzInput15OutputFizzBuzz(t *testing.T) {
	assert.Equal(t, "fizzbuzz", fizzbuzz.FizzBuzz(15))
}

func TestFizzBuzzInput30OutputFizzBuzz(t *testing.T) {
	assert.Equal(t, "fizzbuzz", fizzbuzz.FizzBuzz(30))
}
