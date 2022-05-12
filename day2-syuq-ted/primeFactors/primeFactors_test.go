package primeFactors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnEmpty(t *testing.T) {
	assert.Equal(t, []int{}, primeFactors(1))
}
func TestReturnOne(t *testing.T) {
	assert.Equal(t, []int{2}, primeFactors(2))
	assert.Equal(t, []int{3}, primeFactors(3))
	assert.Equal(t, []int{2}, primeFactors(4))
}
