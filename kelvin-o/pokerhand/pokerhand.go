package pokerhand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPokerHand(t *testing.T) {
	t.Run("it should lose with high A", func(t *testing.T) {
		assert.Equal(t, compareHand("8C TS KC 9H 4S 7D 2S 5D 3S AC"), "lose")
	})
}

func compareHand(s string) string {
	return "lose"
}
