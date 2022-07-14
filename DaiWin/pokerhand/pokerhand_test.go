package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// type PokerHandExamples struct {
// 	WiningHighCard string
// 	LosingHighCard string
// }

const winingHighCard = "2H 3H 4H 5H 8S"
const winingHighCard2 = "2H 3H 4H 5H 9S"
const losingHighCard = "2H 3H 4H 5H 7S"

func TestPlayer1Win0TimesWhenNoGamePlayed(t *testing.T) {
	pokerFile := []string{}
	assert.Equal(t, 0, ValidatePokerFile(pokerFile))
}

func TestPlayer1WinOnceIn1Game(t *testing.T) {
	pokerFile := []string{winingHighCard + " " + losingHighCard}
	assert.Equal(t, 1, ValidatePokerFile(pokerFile))
}
func TestPlayer1WinOnceIn1Game1(t *testing.T) {
	pokerFile := []string{winingHighCard2 + " " + losingHighCard}
	assert.Equal(t, 1, ValidatePokerFile(pokerFile))
}

func TestPlayer1WinOnceIn2Game(t *testing.T) {
	pokerFile := []string{winingHighCard + " " + losingHighCard, losingHighCard + " " + winingHighCard}
	assert.Equal(t, 1, ValidatePokerFile(pokerFile))
}

func TestPlayer1LoseOnceIn1Game(t *testing.T) {
	pokerFile := []string{losingHighCard + " " + winingHighCard}
	assert.Equal(t, 0, ValidatePokerFile(pokerFile))
}

func TestPlayer1LoseOnceIn2Game(t *testing.T) {
	pokerFile := []string{losingHighCard + " " + winingHighCard, winingHighCard + " " + losingHighCard}
	assert.Equal(t, 1, ValidatePokerFile(pokerFile))
}

func TestPlayer1Win2In2Game(t *testing.T) {
	pokerFile := []string{winingHighCard + " " + losingHighCard, winingHighCard + " " + losingHighCard}
	assert.Equal(t, 2, ValidatePokerFile(pokerFile))
}

func TestPlayer1Win2In3Game(t *testing.T) {
	pokerFile := []string{winingHighCard + " " + losingHighCard, winingHighCard + " " + losingHighCard, losingHighCard + " " + winingHighCard}
	assert.Equal(t, 2, ValidatePokerFile(pokerFile))
}

func ValidatePokerFile(list []string) int {
	player1WinningCount := 0
	for _, row := range list {
		if isPlayer1Win(row) {
			player1WinningCount += 1
		}
	}
	return player1WinningCount
}

func isPlayer1Win(row string) bool {
	return row[12] > row[27]
}
