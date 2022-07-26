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

	t.Run("it should say fizz if enter 3", func(t *testing.T) {
		assert.Equal(t, fizzbuzz(3), "fizz")
	})
}

func fizzbuzz(i int) string {
	if i == 3 {
		return "fizz"
	}
	return fmt.Sprintf("%d", i)
}
