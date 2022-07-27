package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FizzBuzz_1_Return_1(t *testing.T) {
	assert.Equal(t, "1", FizzBuzz(1))
}

func xTest_FizzBuzz_3_Return_Fizz(t *testing.T) {
	assert.Equal(t, "fizz", FizzBuzz(3))
}
