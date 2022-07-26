package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("Input 1 get '1'", func(t *testing.T) {
		assert.Equal(t, "1", FizzBuzz(1))
	})

	t.Run("Input 2 get '2'", func(t *testing.T) {
		assert.Equal(t, "2", FizzBuzz(2))
	})

	t.Run("Input 3 get 'fizz'", func(t *testing.T) {
		assert.Equal(t, "fizz", FizzBuzz(3))
	})
}
