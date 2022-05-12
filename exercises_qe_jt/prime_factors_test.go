package exercises_qe_jt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func PrimeFactors(n int) []int {
	result := []int{}
	if n % 2 == 0 {
		// remainder := n / 2
		result = append(result, 2)
		return result
	}
	return append(result, n)
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

func xTestPrimeFactors6(t * testing.T) {
	assert.Equal(t, []int{2, 3}, PrimeFactors(6))
}
