package ivan_sheen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimeFactors1(t *testing.T) {
	assert.Equal(t, []int{}, primeFactors(1))
}

func TestPrimeFactors2(t *testing.T) {
	assert.Equal(t, []int{}, primeFactors(2))
}
