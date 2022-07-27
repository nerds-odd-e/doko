package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("it should return number by default", func(t *testing.T) {
		assert.Equal(t, Fizzbuzz(1), "1")
		assert.Equal(t, Fizzbuzz(2), "2")
	})

	t.Run("it should say fizz if input devisible by 3", func(t *testing.T) {
		assert.Equal(t, Fizzbuzz(3), "fizz")
		assert.Equal(t, Fizzbuzz(6), "fizz")
		assert.Equal(t, Fizzbuzz(9), "fizz")
	})
}
