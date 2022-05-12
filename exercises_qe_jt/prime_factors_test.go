package exercises_qe_jt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func PrimeFactors(n int) string {
	if n % 2 == 0 {
		return "2"
	}
	return fmt.Sprintf("%v", n)
}

func PrimeFactors1(n int) []int {
	if n % 2 == 0 {
		return []int{2}
	}
	return []int{n}
}

func TestPrimeFactors2(t * testing.T) {
	fmt.Println(PrimeFactors1(2))
	assert.Equal(t, []int{2}, PrimeFactors1(2))
}

func TestPrimeFactors3(t * testing.T) {
	assert.Equal(t, []int{3}, PrimeFactors1(3))
}

func TestPrimeFactors4(t * testing.T) {
	assert.Equal(t, []int{2}, PrimeFactors1(4))
}

func TestPrimeFactors5(t * testing.T) {
	assert.Equal(t, "5", PrimeFactors(5))
}

func TestPrimeFactors6(t * testing.T) {
	assert.Equal(t, "2", PrimeFactors(6))
}
