package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrime(t *testing.T) {
	result := findPrime(1)
	assert.Equal(t, []int{}, result)
	result = findPrime(2)
	assert.Equal(t, []int{2}, result)
	result = findPrime(3)
	assert.Equal(t, []int{3}, result)
	result = findPrime(5)
	assert.Equal(t, []int{5}, result)
	result = findPrime(7)
	assert.Equal(t, []int{7}, result)
	result = findPrime(11)
	assert.Equal(t, []int{11}, result)
	result = findPrime(17)
	assert.Equal(t, []int{17}, result)
	result = findPrime(19)
	assert.Equal(t, []int{19}, result)
}

func TestStartEven(t *testing.T) {
	result := findPrime(4)
	assert.Equal(t, []int{2, 2}, result)
	result = findPrime(6)
	assert.Equal(t, []int{2, 3}, result)
	result = findPrime(8)
	assert.Equal(t, []int{2, 2, 2}, result)
	result = findPrime(10)
	assert.Equal(t, []int{2, 5}, result)
	result = findPrime(16)
	assert.Equal(t, []int{2, 2, 2, 2}, result)
	result = findPrime(18)
	assert.Equal(t, []int{2, 3, 3}, result)
	result = findPrime(20)
	assert.Equal(t, []int{2, 2, 5}, result)
}

func TestStartOdd(t *testing.T) {
	result := findPrime(9)
	assert.Equal(t, []int{3, 3}, result)
	result = findPrime(15)
	assert.Equal(t, []int{3, 5}, result)
	result = findPrime(25)
	assert.Equal(t, []int{5, 5}, result)
	result = findPrime(49)
	assert.Equal(t, []int{7, 7}, result)
}

func xTestStartOdds(t *testing.T) {

	result := findPrime(-1)
	assert.Equal(t, []int{}, result)
}

func findPrime(num int) []int {
	if num == 1 {
		return []int{}
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return append([]int{i}, findPrime(num/i)...)
		}
	}
	return []int{num}
}
