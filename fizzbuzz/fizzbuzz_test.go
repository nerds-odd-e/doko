package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// - should test for number divisible by 3 ONLY, print FIZZ
// - should test for number divisible by 5 ONLY, print BUZZ
// - should test for number divisible by 3 and 5, print FIZZBUZZ
// - should test for any other numbers, print number itself

func TestForDivisible3(t *testing.T) {
	assert.Equal(t, "FIZZ", fizzbuzz(3))
	assert.Equal(t, "FIZZ", fizzbuzz(6))
}

func TestForDivisible5(t *testing.T) {
	assert.Equal(t, "BUZZ", fizzbuzz(5))
}
