package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizz(t *testing.T) {
	result := findFizzBuzz(3)
	assert.Equal(t, "Fizz", result)
	result = findFizzBuzz(6)
	assert.Equal(t, "Fizz", result)
	result = findFizzBuzz(99)
	assert.Equal(t, "Fizz", result)
}

func TestBuzz(t *testing.T) {
	result := findFizzBuzz(5)
	assert.Equal(t, "Buzz", result)
	result = findFizzBuzz(10)
	assert.Equal(t, "Buzz", result)
}

func TestFizzBuzz(t *testing.T) {
	result := findFizzBuzz(15)
	assert.Equal(t, "FizzBuzz", result)
	result = findFizzBuzz(30)
	assert.Equal(t, "FizzBuzz", result)
}

func TestNotFizzBuzz(t *testing.T) {
	result := findFizzBuzz(1)
	assert.Equal(t, "1", result)
	result = findFizzBuzz(2)
	assert.Equal(t, "2", result)
	result = findFizzBuzz(4)
	assert.Equal(t, "4", result)
	result = findFizzBuzz(7)
	assert.Equal(t, "7", result)
	result = findFizzBuzz(8)
	assert.Equal(t, "8", result)

}

func findFizzBuzz(num int) string {
	if num%15 == 0 {
		return "FizzBuzz"
	}
	if num%5 == 0 {
		return "Buzz"
	}
	if num%3 == 0 {
		return "Fizz"
	}
	return fmt.Sprintf("%v", num)
}
