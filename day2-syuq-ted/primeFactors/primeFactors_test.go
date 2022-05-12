package primeFactors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroPrimeFactors(t *testing.T) {
	assert.Equal(t, []int{2}, primeFactors(2))
}
