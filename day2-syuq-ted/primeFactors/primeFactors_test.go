package primeFactors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroPrimeFactors(t *testing.T) {
	assert.Equal(t, []int{}, primeFactors(2))
}
