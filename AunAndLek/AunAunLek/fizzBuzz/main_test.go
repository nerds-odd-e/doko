package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("Input 1 get empty string", func(t *testing.T) {
		assert.Equal(t, "", FizzBuzz(1))
	})
}
