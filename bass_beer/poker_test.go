package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer1Win(t *testing.T) {
	assert.Equal(t, true, Player1Win("8C TS KC 9H 4S 7D 2S 5D 3S QC")) // in case of A win
}

func TestPlayer1Lose(t *testing.T) {
	assert.Equal(t, false, Player1Win("8C TS KC 9H 4S 7D 2S 5D 3S AC"))
}

func TestPlayer1LoseWithAce(t *testing.T) {
	assert.Equal(t, false, Player1Win("8C TS KC 9H 4S 7D 2S 5D 3S AS"))

}

func TestPlayer1LoseWithKing(t *testing.T) {
	assert.Equal(t, false, Player1Win("8C TS QC 9H 4S 7D 2S 5D 3S KS"))
}

func TestPlayer1LoseWithKingOnNewPosition(t *testing.T) {
	assert.Equal(t, false, Player1Win("8C TS QC 9H 4S 7D 2S 5D KS 3S"))
}

func TestPlayer1LoseWithJ(t *testing.T) {
	assert.Equal(t, false, Player1Win("8C TS 2C 9H 4S 7D 3S JS 5D 2S"))
}

func TestPlayer1WinHighCard(t *testing.T) {
	assert.Equal(t, true, Player1Win("8C TS 2C 9H 4S 7D 3S 9S 5D 2S"))
}

func TestPlayer1WinPercentage(t *testing.T) {
	assert.Equal(t, 100.0, GetPlayer1Winrate([]string{"8C TS 2C 9H 4S 7D 3S 9S 5D 2S"}))
}

func TestPlayer1WinZeroPercentage(t *testing.T) {
	assert.Equal(t, 0.0, GetPlayer1Winrate([]string{"8C TS 2C 9H 4S 7D 3S JS 5D 2S"}))
}

func TestEmptyGameZeroPercentage(t *testing.T) {
	assert.Equal(t, 0.0, GetPlayer1Winrate([]string{}))
}

func TestTwoGamesPlayer_win100Percentage(t *testing.T) {
	games := []string {
		"8C TS 2C 9H 4S 7D 3S 9S 5D 2S",
		"8C TS 2C 9H 4S 7D 3S 9S 5D 2S",
	}
	assert.Equal(t, 100.0, GetPlayer1Winrate(games))
}