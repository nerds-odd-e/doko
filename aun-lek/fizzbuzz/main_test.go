package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("Input 1 get '1'", func(t *testing.T) {
		assert.Equal(t, "1", FizzBuzz(1))
	})
}
