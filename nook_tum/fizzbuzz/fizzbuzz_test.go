package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInput1_ShouldSay1(t *testing.T) {
	answer := FizzBuzz(1)
	expected := ""
	assert.Equal(t, answer, expected, "")
}
