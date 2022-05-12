package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfirmFail(t *testing.T) {
	assert.Equal(t, true, true)
}

func TestOneReturnOne(t *testing.T) {
	assert.Equal(t, "1", fizzbuzz(1))
	assert.Equal(t, "1", fizzbuzz(2))
}
