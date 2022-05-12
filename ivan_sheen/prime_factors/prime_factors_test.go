package ivan_sheen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimeFactors1(t *testing.T) {
	assert.Equal(t, []int{}, primeFactors(1))
}

func TestPrimeFactors2(t *testing.T) {
	assert.Equal(t, []int{2}, primeFactors(2))
}

func TestPrimeFactors3(t *testing.T) {
	assert.Equal(t, []int{3}, primeFactors(3))
}

func TestPrimeFactors4(t *testing.T) {
	assert.Equal(t, []int{2}, primeFactors(4))
}

func TestPrimeFactors5(t *testing.T) {
	assert.Equal(t, []int{2}, primeFactors(4))
}
