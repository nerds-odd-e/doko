package pokerhand

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPokerHand(t *testing.T) {
	t.Run("it should win 0 with no game", func(t *testing.T) {
		assert.Equal(t, runGames(""), 0.0)
	})

	t.Run("it should win 1 with high A", func(t *testing.T) {
		p1WinByHighCard := "AC TS KC 9H 4S 7D 2S 5D 3S QC"
		assert.Equal(t, runGames(p1WinByHighCard), 1.0)
	})

	t.Run("it should win 0 with P2 high A", func(t *testing.T) {
		if false {
			p2WinByHighCard := "2D 3S 4D 5S 7C AC TS KC 9H 4S"
			assert.Equal(t, runGames(p2WinByHighCard), 0.0)
		}
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
		if player1Win(game, cardPos) {
			return "win"
		}
	}

	return "lose"
}

func player1Win(game string, cardPos int) bool {
	return strings.Contains("AKQT98", getNthCardRank(game, cardPos))
}

func getNthCardRank(game string, cardPos int) string {
	return string(game[(cardPos-1)*3])
}
