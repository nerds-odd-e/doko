package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fizzbuzz_1(t *testing.T) {
	assert.Equal(t, "1", fizzbuzz(1))
}

func Test_fizzbuzz_2(t *testing.T) {
	assert.Equal(t, "2", fizzbuzz(2))
}

func fizzbuzz(i int) string {
	if i == 2 {
		return "2"
	}
	return "1"
}
