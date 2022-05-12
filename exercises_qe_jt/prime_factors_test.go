package exercises_qe_jt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func PrimeFactors(n int) []int {
	if n % 2 == 0 {
		return []int{2}
	}
	return []int{n}
}

func TestPrimeFactors2(t * testing.T) {
	assert.Equal(t, []int{2}, PrimeFactors(2))
}

func TestPrimeFactors3(t * testing.T) {
	assert.Equal(t, []int{3}, PrimeFactors(3))
}

func TestPrimeFactors4(t * testing.T) {
	assert.Equal(t, []int{2}, PrimeFactors(4))
}

func TestPrimeFactors5(t * testing.T) {
	assert.Equal(t, []int{5}, PrimeFactors(5))
}

func TestPrimeFactors6(t * testing.T) {
	assert.Equal(t, []int{2}, PrimeFactors(6))
}
