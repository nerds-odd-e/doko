package exercises_qe_jt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func PrimeFactors(n int) string {
	return "2"
}

func TestPrimeFactors2(t * testing.T) {
	assert.Equal(t, "2", PrimeFactors(2))
}