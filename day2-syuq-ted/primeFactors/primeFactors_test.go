package primeFactors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnePrimeFactors(t *testing.T) {
	assert.Equal(t, []int{}, primeFactors(1))

	assert.Equal(t, []int{2}, primeFactors(2))
	assert.Equal(t, []int{3}, primeFactors(3))
	assert.Equal(t, []int{5}, primeFactors(5))

	// squares
	assert.Equal(t, []int{2}, primeFactors(4))
	assert.Equal(t, []int{3}, primeFactors(9))
	assert.Equal(t, []int{5}, primeFactors(25))
}

func TestTwoPrimeFactors(t *testing.T) {
	assert.Equal(t, []int{2, 3}, primeFactors(6))
}
