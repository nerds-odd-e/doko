package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeReturnFizz(t *testing.T) {
	assert.Equal(t, "Fizz", FizzBuzz(3))
}
