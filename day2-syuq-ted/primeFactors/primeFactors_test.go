package primeFactors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneReturnEmpty(t *testing.T) {
	assert.Equal(t, []int{}, primeFactors(1))
}
func TestOneReturnOne(t *testing.T) {
	assert.Equal(t, []int{2}, primeFactors(2))
}
