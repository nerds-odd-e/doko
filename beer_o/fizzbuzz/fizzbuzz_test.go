package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("it should return number by default", func(t *testing.T) {
		for _, input := range []int{1, 2, 4} {
			assert.Equal(t, Fizzbuzz(input), strconv.Itoa(input))
		}
	})

	t.Run("it should say fizz if input devisible by 3", func(t *testing.T) {
		for _, input := range []int{3, 6, 9} {
			assert.Equal(t, Fizzbuzz(input), "fizz")
		}
	})
}
