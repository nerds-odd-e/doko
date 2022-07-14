package winmeen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumber1(t *testing.T) {
	assert.Equal(t, Fizzbuzz(0), "")
}
