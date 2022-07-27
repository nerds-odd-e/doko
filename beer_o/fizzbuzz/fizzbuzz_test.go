package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("it should say 1 if enter 1", func(t *testing.T) {
		assert.Equal(t, Fizzbuzz(1), "1")
	})

	t.Run("it should say 2 if enter 2", func(t *testing.T) {
		assert.Equal(t, Fizzbuzz(2), "2")
	})

	t.Run("it should say fizz if enter 3", func(t *testing.T) {
		assert.Equal(t, Fizzbuzz(3), "fizz")
	})

	t.Run("it should say fizz if enter 6", func(t *testing.T) {
		assert.Equal(t, Fizzbuzz(6), "fizz")
	})
}
