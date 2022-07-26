package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz1(t *testing.T) {
	if fizzbuzz(1) != "1" {
		t.Errorf("Expected %q, but got %q", "1", fizzbuzz(1))
	}
}

func TestFizzBuzz2(t *testing.T) {
	assert.Equal(t, fizzbuzz(1), "1")
}

func fizzbuzz(i int) string {
	return "1"
}
