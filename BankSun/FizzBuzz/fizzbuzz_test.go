package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_InputFizzbuzz(t *testing.T) {
	a := inputFizzbuzz(1)
	b := "1"
	assert.Equal(t, a, b)
}

func Test_InputFizzbuzzDivisible3(t *testing.T) {
	a := inputFizzbuzz(3)
	b := "Fizz"
	assert.Equal(t, a, b)
}

func Test_InputFizzbuzzDivisible5(t *testing.T) {
	a := inputFizzbuzz(5)
	b := "Buzz"
	assert.Equal(t, a, b)
}

func Test_InputFizzbuzzDivisible35(t *testing.T) {
	a := inputFizzbuzz(15)
	b := "FizzBuzz"
	assert.Equal(t, a, b)
}

func Test_InputFizzbuzzByValue(t *testing.T) {
	a := inputFizzbuzz(14)
	b := "14"
	assert.Equal(t, a, b)
}
