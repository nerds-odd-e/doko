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
		assert.Equal(t, compareHand("TS AC KC 9H 4S 7D 2S 5D 3S QC"), "lose")
	})
}

func compareHand(game string) string {
	p1FirstCardRank := string(game[0])
	if p1FirstCardRank == "A" {
		return "win"
	}

	return "lose"
}
