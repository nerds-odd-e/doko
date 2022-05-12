package primeFactors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnePrimeFactors(t *testing.T) {
	assert.Equal(t, []int{}, primeFactors(1))
}
