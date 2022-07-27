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

func xTestFizzBuzzInput3OutputFizz(t *testing.T) {
	assert.Equal(t, "fizz", fizzbuzz.FizzBuzz(3))
}
