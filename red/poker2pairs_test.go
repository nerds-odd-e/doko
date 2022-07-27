package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoPairs(t *testing.T) {
	t.Run("Player1 should win two pairs at rate 1.0", func(t *testing.T) {
		games := []string{"AC AS KC KS 2C 3C 3S 6C 6S 7S"}
		assert.Equal(t, WinRateFor2Pairs(games), 1.0)
	})
}
