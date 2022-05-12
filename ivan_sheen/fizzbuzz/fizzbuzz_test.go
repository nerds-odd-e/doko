package ivan_sheen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz1(t *testing.T) {
	assert.Equal(t, "1", fizzbuzz(1))
}

func TestFizzBuzz2(t *testing.T) {
	assert.Equal(t, "2", fizzbuzz(2))
}

func TestFizzBuzz4(t *testing.T) {
	assert.Equal(t, "4", fizzbuzz(4))
}

func TestFizzBuzz3(t *testing.T) {
	assert.Equal(t, "fizz", fizzbuzz(3))
}

func TestFizzBuzz6(t *testing.T) {
	assert.Equal(t, "fizz", fizzbuzz(6))
}

func TestFizzBuzz5(t *testing.T) {
	assert.Equal(t, "buzz", fizzbuzz(5))
}

func TestFizzBuzz10(t *testing.T) {
	assert.Equal(t, "buzz", fizzbuzz(10))
}

func TestFizzBuzz15(t *testing.T) {
	assert.Equal(t, "buzz", fizzbuzz(15))
}
