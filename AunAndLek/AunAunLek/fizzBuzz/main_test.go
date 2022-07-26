package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {

	expected := "1"
	input := 1
	t.Run(fmt.Sprintf("Input 1 get '%v'", expected), func(t *testing.T) {
		assert.Equal(t, expected, FizzBuzz(input))
	})

	t.Run("Input 2 get '2'", func(t *testing.T) {
		assert.Equal(t, "2", FizzBuzz(2))
	})

	t.Run("Input 3 get 'fizz'", func(t *testing.T) {
		assert.Equal(t, "fizz", FizzBuzz(3))
	})

	t.Run("Input 5 get 'buzz'", func(t *testing.T) {
		assert.Equal(t, "buzz", FizzBuzz(5))
	})

	t.Run("Input 6 get 'fizz'", func(t *testing.T) {
		assert.Equal(t, "fizz", FizzBuzz(6))
	})

	t.Run("Input 10 get 'buzz'", func(t *testing.T) {
		assert.Equal(t, "buzz", FizzBuzz(10))
	})
}
