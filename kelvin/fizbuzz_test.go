package fizzbuzz

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz1(t *testing.T) {
	assert.Equal(t, fizzbuzz(1), "1")
}

func TestFizzBuzz2(t *testing.T) {
	assert.Equal(t, fizzbuzz(2), "2")
}

func fizzbuzz(i int) string {
	return fmt.Sprintf("%d", i)
}
