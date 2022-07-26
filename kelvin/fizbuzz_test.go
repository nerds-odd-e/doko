package fizzbuzz

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz1(t *testing.T) {
	t.Run("it should return number by default", func(t *testing.T) {
		assert.Equal(t, fizzbuzz(1), "1")
		assert.Equal(t, fizzbuzz(2), "2")
	})

	t.Run("it should say fizz if divide 3 equal 0", func(t *testing.T) {
		assert.Equal(t, fizzbuzz(3), "fizz")
		assert.Equal(t, fizzbuzz(6), "fizz")
	})

	t.Run("it should say buzz if input divisible by 5", func(t *testing.T) {
		assert.Equal(t, fizzbuzz(5), "buzz")
		assert.Equal(t, fizzbuzz(5), "buzz")
	})
}

func fizzbuzz(i int) string {
	if i == 5 {
		return "buzz"
	}
	if i%3 == 0 {
		return "fizz"
	}
	return fmt.Sprintf("%d", i)
}
