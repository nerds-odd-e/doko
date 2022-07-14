package winmeen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumber1(t *testing.T) {
	assert.Equal(t, Fizzbuzz(1), "1")
}

func TestNumber3(t *testing.T) {
	assert.Equal(t, Fizzbuzz(3), "Fizz")
}
func TestNumber5(t *testing.T) {
	assert.Equal(t, Fizzbuzz(5), "Buzz")
}

func TestNumber6(t *testing.T) {
	assert.Equal(t, Fizzbuzz(6), "Fizz")
}
func TestNumber15(t *testing.T) {
	assert.Equal(t, Fizzbuzz(15), "FizzBuzz")
}

func TestNumber16(t *testing.T) {
	assert.Equal(t, Fizzbuzz(16), "1")
}
