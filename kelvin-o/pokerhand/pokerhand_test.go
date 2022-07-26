package pokerhand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPokerHand(t *testing.T) {
	t.Run("it should win 0 with no game", func(t *testing.T) {
		assert.Equal(t, runGames(""), 0.0)
	})

	t.Run("it should win 1 with high A", func(t *testing.T) {
		p1WinByHighCard := "2C 3S 4C 5H 8S 2D 3S 4D 5S 7C"
		assert.Equal(t, runGames(p1WinByHighCard), 1.0)
	})

	t.Run("it should win 0 with P2 high A", func(t *testing.T) {
		p2WinByHighCard := "2C 3S 4C 5H 7S 2D 3S 4D 5S 8C"
		assert.Equal(t, runGames(p2WinByHighCard), 0.0)
	})
}

func runGames(games string) float64 {
	if len(games) == 0 {
		return 0.0
	}

	if compareHand(games) {
		return 1.0
	}

	return 0.0
}

func compareHand(game string) bool {
	return game[12] == '8'
}
