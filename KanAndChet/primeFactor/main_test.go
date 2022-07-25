package primefactor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputOne(t *testing.T) {
	result := primeFactor(1)
	assert.Equal(t, len(result), 0, "1 is not prime number so we should get empty array")
}

func TestInputTwo(t *testing.T) {
	result := primeFactor(2)
	assert.Equal(t, result, []int{2}, "2 is prime number so we prefer to get [2]")
}

func TestInputThree(t *testing.T) {
	result := primeFactor(3)
	assert.Equal(t, result, []int{3}, "3 is prime number so we prefer to get [3]")
}

func TestInputFour(t *testing.T) {
	result := primeFactor(4)
	assert.Equal(t, result, []int{2, 2}, "4 is not a prime number so we should get [2, 2]")
}

func TestInputFive(t *testing.T) {
	result := primeFactor(5)
	assert.Equal(t, result, []int{5}, "5 is prime number so we prefer to get [5]")
}

func TestInputSix(t *testing.T) {
	result := primeFactor(6)
	assert.Equal(t, result, []int{2, 3}, "6 is not a prime number so we should get [2, 3]")
}

func TestInputNine(t *testing.T) {
	result := primeFactor(9)
	assert.Equal(t, result, []int{3, 3}, "9 is not a prime number so we should get [3, 3]")
}
