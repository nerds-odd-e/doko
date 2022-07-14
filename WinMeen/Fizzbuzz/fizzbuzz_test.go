package winmeen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumber0(t *testing.T) {
	assert.Equal(t, Fizzbuzz(0), "0")
}

func TestNumber1(t *testing.T) {
	assert.Equal(t, Fizzbuzz(1), "1")
}