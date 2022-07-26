package pokerhand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPokerHand(t *testing.T) {
	t.Run("it should win 0", func(t *testing.T) {
		assert.Equal(t, runGames(""), 0.0)
		assert.Equal(t, runGames("8C TS KC 9H 4S 7D 2S 5D 3S AC"), 0.0)
	})

	t.Run("it should win 1 with high A", func(t *testing.T) {
		assert.Equal(t, runGames("AC TS KC 9H 4S 7D 2S 5D 3S QC"), 1.0)
	})

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

func runGames(games string) float64 {
	if len(games) == 0 {
		return 0.0
	}

	if compareHand(games) == "win" {
		return 1.0
	}

	return 0.0
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
