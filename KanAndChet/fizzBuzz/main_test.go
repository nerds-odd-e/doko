package AunAndChet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputNumber(t *testing.T) {
	result, _ := fizzBuzz(1)
	assert.Equal(t, result, "1", "Invalid output, output not string")
}

func TestFizz(t *testing.T) {
	result, _ := fizzBuzz(3)
	assert.Equal(t, result, "Fizz", "Input is multiply by 3, so we should get Fizz")
}

func TestBuzz(t *testing.T) {
	result, _ := fizzBuzz(5)
	assert.Equal(t, result, "Buzz", "Input is multiply by 5, so we should get Buzz")
}

func TestInputTen(t *testing.T) {
	errMsg := "Ten divide by 5, so we should get Buzz"
	result, _ := fizzBuzz(10)
	assert.Equal(t, result, "Buzz", errMsg)
}

func TestInputSix(t *testing.T) {
	errMsg := "Six divide by 3, so we should get Fizz"
	result, _ := fizzBuzz(6)
	assert.Equal(t, result, "Fizz", errMsg)
}

func TestInputFifteen(t *testing.T) {
	result, _ := fizzBuzz(15)
	assert.Equal(t, result, "FizzBuzz", "Input : 15, so we should get FizzBuzz")
}

func TestInputThirdTy(t *testing.T) {
	result, _ := fizzBuzz(30)
	assert.Equal(t, result, "FizzBuzz", "Input : 30, so we should get FizzBuzz")
}
