package primeFactors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnePrimeFactors(t *testing.T) {
	assert.Equal(t, []int{2}, primeFactors(2))
	assert.Equal(t, []int{3}, primeFactors(3))
}
