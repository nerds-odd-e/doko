package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeReturnFizz(t *testing.T) {
	assert.Equal(t, "", FizzBuzz(3))
}
