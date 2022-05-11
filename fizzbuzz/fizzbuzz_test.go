package fizzbuzz

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnFizzIfDivisibleByThree(t *testing.T) {
	//  Arrange
	num := 3
	// Act
	output := FizzBuzz(num)

	// Assert
	assert.Equal(t, "Fizz", output)
	assert.Equal(t, "Fizz", FizzBuzz(6))
	assert.Equal(t, "Fizz", FizzBuzz(9))
	assert.Equal(t, "Fizz", FizzBuzz(12))
}

func TestReturnNumberIfNotDivisibleByThree(t *testing.T) {
	//  Arrange
	num := 2
	// Act
	output := FizzBuzz(num)

	// Assert
	assert.Equal(t, fmt.Sprint(num), output)
	assert.Equal(t, fmt.Sprint(1), FizzBuzz(1))
	assert.Equal(t, fmt.Sprint(4), FizzBuzz(4))
	assert.Equal(t, fmt.Sprint(7), FizzBuzz(7))
}

func TestReturnBuzzIfDivisibleByFive(t *testing.T) {
	//  Arrange
	num := 5
	// Act
	output := FizzBuzz(num)

	// Assert
	assert.Equal(t, "Buzz", output)
	assert.Equal(t, "Buzz", FizzBuzz(10))
	assert.Equal(t, "Buzz", FizzBuzz(20))
	assert.Equal(t, "Buzz", FizzBuzz(25))
}

func TestReturnFizzBuzzIfDivisibleByThreeAndFive(t *testing.T) {
	//  Arrange
	num := 15
	// Act
	output := FizzBuzz(num)

	// Assert
	assert.Equal(t, "FizzBuzz", output)
	assert.Equal(t, "FizzBuzz", FizzBuzz(30))
	assert.Equal(t, "FizzBuzz", FizzBuzz(45))
}
