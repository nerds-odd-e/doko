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
}

func fizzbuzz(i int) string {
	return fmt.Sprintf("%d", i)
}
