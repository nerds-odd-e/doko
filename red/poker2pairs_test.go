package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoPairs(t *testing.T) {
	t.Run("Player1 should win two pairs at rate 1.0", func(t *testing.T) {
		cards := []string{"AC", "AS", "KC", "KS", "2C"}
		assert.Equal(t, IsTwoPair(cards), true)
	})
}
