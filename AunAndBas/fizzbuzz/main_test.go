package fizzbuzz

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	testcase := map[int]string{1: "1",2: "2"}
	for input, expected := range testcase {
		t.Run(fmt.Sprintf("Input %d should get '%s'", input, expected), func(t *testing.T) {
			// When
			output := FizzBuzz(input)

			// Then
			assert.Equal(t, expected, output)
		})
	}

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
