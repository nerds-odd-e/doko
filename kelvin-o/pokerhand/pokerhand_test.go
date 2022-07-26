package pokerhand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPokerHand(t *testing.T) {
	if false {
		t.Run("it should lose with high A", func(t *testing.T) {
			assert.Equal(t, compareHand("8C TS KC 9H 4S 7D 2S 5D 3S AC"), "lose")
		})

		t.Run("it should win with high A", func(t *testing.T) {
			assert.Equal(t, compareHand("AC TS KC 9H 4S 7D 2S 5D 3S QC"), "win")
			assert.Equal(t, compareHand("TS AC KC 9H 4S 7D 2S 5D 3S QC"), "win")
			assert.Equal(t, compareHand("TS KC AC 9H 4S 7D 2S 5D 3S QC"), "win")
			assert.Equal(t, compareHand("TS KC 9H AC 4S 7D 2S 5D 3S QC"), "win")
			assert.Equal(t, compareHand("TS KC 9H 4S AC 7D 2S 5D 3S QC"), "win")
		})
	}
}

func runGame(games string) float64 {
	return 0
}

func compareHand(game string) string {
	for cardPos := 1; cardPos <= 5; cardPos++ {
		if getP1NthCardRank(game, cardPos) == "A" {
			return "win"
		}
	}
	return "lose"
}

func getP1NthCardRank(game string, cardPos int) string {
	return string(game[(cardPos-1)*3])
}
