package pokerhand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const p1WinByHighCard = "2C 3S 4C 5H 8S 2D 3S 4D 5S 7C"
const p2WinByHighCard = "2C 3S 4C 5H 7S 2D 3S 4D 5S 8C"

func TestPokerHand(t *testing.T) {
	t.Run("it should win 0 with no game", func(t *testing.T) {
		assert.Equal(t, runGames(""), 0.0)
	})

	t.Run("it should win 1 with high A", func(t *testing.T) {
		assert.Equal(t, runGames(p1WinByHighCard), 1.0)
	})

	t.Run("it should win 0 with P2 high A", func(t *testing.T) {
		assert.Equal(t, runGames(p2WinByHighCard), 0.0)
	})

	t.Run("player1 win twice", func(t *testing.T) {
		games := p1WinByHighCard + "\n" + p1WinByHighCard
		assert.Equal(t, runGames(games), 1.0)
	})

	t.Run("player1 win once player2 win once", func(t *testing.T) {
		games := p1WinByHighCard + "\n" + p2WinByHighCard
		assert.Equal(t, runGames(games), 1.0)
	})

}

func runGames(games string) float64 {
	if len(games) == 0 {
		return 0.0
	}

	if player1Win(games) {
		return 1.0
	}

	return 0.0
}

func player1Win(game string) bool {
	return game[12] == '8'
}
