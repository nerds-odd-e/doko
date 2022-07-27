package chet_kelvin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInput1(t *testing.T) {
	assert.Equal(t, "1", fizzbuzz(1))
}
