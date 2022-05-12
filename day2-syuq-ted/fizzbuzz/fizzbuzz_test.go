package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfirmFail(t *testing.T) {
	assert.Equal(t, true, true)
}

func TestNormalReturnNormal(t *testing.T) {
	assert.Equal(t, "1", fizzbuzz(1))
	assert.Equal(t, "2", fizzbuzz(2))
}

func TestDivisible3(t *testing.T) {
	assert.Equal(t, "FIZZ", fizzbuzz(3))
	assert.Equal(t, "FIZZ", fizzbuzz(6))
}

func TestDivisible5(t *testing.T) {
	assert.Equal(t, "BUZZ", fizzbuzz(5))
	assert.Equal(t, "BUZZ", fizzbuzz(10))
}

func TestDivisible15(t *testing.T) {
	assert.Equal(t, "FIZZBUZZ", fizzbuzz(15))
}
