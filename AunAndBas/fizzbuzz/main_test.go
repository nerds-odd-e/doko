package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("Input 1 should get '1'", func(t *testing.T) {
		// Given
		input:= 1

		// When
		output := FizzBuzz(input)

		// Then
		assert.Equal(t, output, "1")
	})
}
