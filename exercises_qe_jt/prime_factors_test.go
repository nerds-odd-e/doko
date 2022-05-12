package exercises_qe_jt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func PrimeFactors(n int) string {
	return fmt.Sprintf("%v", n)
}

func TestPrimeFactors2(t * testing.T) {
	assert.Equal(t, "2", PrimeFactors(2))
}

func TestPrimeFactors3(t * testing.T) {
	assert.Equal(t, "3", PrimeFactors(3))
}

func TestPrimeFactors4(t * testing.T) {
	assert.Equal(t, "", PrimeFactors(4))
}

func TestPrimeFactors5(t * testing.T) {
	assert.Equal(t, "5", PrimeFactors(5))
}

func xTestPrimeFactors6(t * testing.T) {
	assert.Equal(t, "6", PrimeFactors(6))
}
