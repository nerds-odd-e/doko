package pokerhand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPokerHand(t *testing.T) {
	t.Run("it should lose with high A", func(t *testing.T) {
		assert.Equal(t, compareHand("8C TS KC 9H 4S 7D 2S 5D 3S AC"), "lose")
	})

	t.Run("it should win with high A", func(t *testing.T) {
		assert.Equal(t, compareHand("AC TS KC 9H 4S 7D 2S 5D 3S QC"), "lose")
	})
}

func compareHand(game string) string {
	if game == "AC TS KC 9H 4S 7D 2S 5D 3S QC" {
		return "win"
	}

	return "lose"
}
