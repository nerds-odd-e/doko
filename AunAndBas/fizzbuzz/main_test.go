package fizzbuzz

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	testcase := map[int]string{1: "1",}
	t.Run(fmt.Sprintf("Input 1 should get '%s'", testcase[1]), func(t *testing.T) {
		// Given
		input:= 1
		expected:= testcase[input]

		// When
		output := FizzBuzz(input)

		// Then
		assert.Equal(t, expected, output)
	})

	t.Run("Input 2 should get '2'", func(t *testing.T) {
		// Given
		input:= 2
		expected:= "2"

		// When
		output := FizzBuzz(input)

		// Then
		assert.Equal(t, expected, output)
	})

	t.Run("Input 3 should get 'fizz'", func(t *testing.T) {
		// Given
		input:= 3
		expected:= "fizz"

		// When
		output := FizzBuzz(input)

		// Then
		assert.Equal(t, expected, output)
	})
}
