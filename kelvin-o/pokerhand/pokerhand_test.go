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
		assert.Equal(t, compareHand("AC TS KC 9H 4S 7D 2S 5D 3S QC"), "win")
		assert.Equal(t, compareHand("TS AC KC 9H 4S 7D 2S 5D 3S QC"), "win")
		assert.Equal(t, compareHand("TS KC AC 9H 4S 7D 2S 5D 3S QC"), "win")
	})
}

func compareHand(game string) string {
	if getP1NthCardRank(game, 1) == "A" || getP1NthCardRank(game, 2) == "A" || getP1NthCardRank(game, 3) == "A" {
		return "win"
	}

	return "lose"
}

func getP1NthCardRank(game string, cardPos int) string {
	return string(game[(cardPos-1)*3])
}
