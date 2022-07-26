package pokerhand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPokerHand(t *testing.T) {
	t.Run("it should win 0", func(t *testing.T) {
		assert.Equal(t, runGames(""), 0.0)
		assert.Equal(t, runGames("2D 3S 4D 5S 7C AC TS KC 9H 4S"), 0.0)
	})

	t.Run("it should win 1 with high A", func(t *testing.T) {
		assert.Equal(t, runGames("AC TS KC 9H 4S 7D 2S 5D 3S QC"), 1.0)
		assert.Equal(t, runGames("TS AC KC 9H 4S 7D 2S 5D 3S QC"), 1.0)
		assert.Equal(t, runGames("TS KC AC 9H 4S 7D 2S 5D 3S QC"), 1.0)
		assert.Equal(t, runGames("TS KC 9H AC 4S 7D 2S 5D 3S QC"), 1.0)
		assert.Equal(t, runGames("TS KC 9H 4S AC 7D 2S 5D 3S QC"), 1.0)
	})

	t.Run("it should win 1 with high K", func(t *testing.T) {
		assert.Equal(t, runGames("2C TS KC 9H 4S 7D 2S 5D 3S QC"), 1.0)
	})

	t.Run("it should win 1 with high Q", func(t *testing.T) {
		assert.Equal(t, runGames("2C TS QC 9H 4S 7D 2S 5D 3S 4C"), 1.0)
	})

	t.Run("it should win 1 with high T", func(t *testing.T) {
		assert.Equal(t, runGames("2C TS 8C 9H 4S 7D 2S 5D 3S 4C"), 1.0)
	})

	t.Run("it should win 1 with high 9", func(t *testing.T) {
		assert.Equal(t, runGames("2C 7S 8C 9H 4S 7D 2S 5D 3S 4C"), 1.0)
	})

	t.Run("it should win 1 with high 8", func(t *testing.T) {
		assert.Equal(t, runGames("2C 7S 8C 6H 4S 7D 2S 5D 3S 4C"), 1.0)
	})
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
		if getP1NthCardRank(game, cardPos) == "K" {
			return "win"
		}
		if getP1NthCardRank(game, cardPos) == "Q" {
			return "win"
		}
		if getP1NthCardRank(game, cardPos) == "T" {
			return "win"
		}
		if getP1NthCardRank(game, cardPos) == "9" {
			return "win"
		}
		if getP1NthCardRank(game, cardPos) == "8" {
			return "win"
		}
	}
	return "lose"
}

func getP1NthCardRank(game string, cardPos int) string {
	return string(game[(cardPos-1)*3])
}
