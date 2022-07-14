package models

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// type PokerHandExamples struct {
// 	WiningHighCard string
// 	LosingHighCard string
// }

const winingHighCard = "2H 3H 4H 5H 8S"
const winingHighCard2 = "2H 3H 4H 5H 9S"
const winingHighCard3 = "2H 3H 4H 5H KS"
const winingHighCard4 = "2H 3H 4H 5H AS"
const winingHighCard5 = "2H 3H 4H 5H QS"
const winingHighCard6 = "2H 3H 4H 5H JS"
const winingHighCard7 = "2H 3H 4H 5H TS"
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
	pokerFile := []string{winingHighCard + " " + losingHighCard, winingHighCard + " " + losingHighCard, losingHighCard + " " + losingHighCard}
	assert.Equal(t, 2, ValidatePokerFile(pokerFile))
}

func TestPlayer1WinOnceUn1GameWithKingCard(t *testing.T) {
	pokerFile := []string{winingHighCard3 + " " + losingHighCard}
	assert.Equal(t, 1, ValidatePokerFile(pokerFile))
}

func TestPlayer1WinOnceUn1GameWithAceCard(t *testing.T) {
	pokerFile := []string{winingHighCard4 + " " + winingHighCard3}
	assert.Equal(t, 1, ValidatePokerFile(pokerFile))
}

func TestPlayer1WinOnceUn1GameWithQueenCard(t *testing.T) {
	pokerFile := []string{winingHighCard5 + " " + winingHighCard6}
	assert.Equal(t, 1, ValidatePokerFile(pokerFile))
}

func TestPlayer1WinOnceUn1GameWithJackCard(t *testing.T) {
	pokerFile := []string{winingHighCard6 + " " + winingHighCard7}
	assert.Equal(t, 1, ValidatePokerFile(pokerFile))
}

func TestPlayer1WinOnceUn1GameWithTenCard(t *testing.T) {
	pokerFile := []string{winingHighCard7 + " " + winingHighCard2}
	assert.Equal(t, 1, ValidatePokerFile(pokerFile))
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
	covertP1Card := string(row[12])
	covertP2Card := string(row[27])
	if row[12] == 'K' {
		covertP1Card = "13"
	}

	if row[12] == 'A' {
		covertP1Card = "14"
	}

	if row[12] == 'Q' {
		covertP1Card = "12"
	}

	if row[12] == 'J' {
		covertP1Card = "11"
	}

	if row[12] == 'T' {
		covertP1Card = "10"
	}

	if row[27] == 'K' {
		covertP2Card = "13"
	}

	if row[27] == 'A' {
		covertP2Card = "14"
	}

	if row[27] == 'Q' {
		covertP2Card = "12"
	}

	if row[27] == 'J' {
		covertP2Card = "11"
	}

	if row[27] == 'T' {
		covertP2Card = "10"
	}

	p1CardPoint, err := strconv.Atoi(covertP1Card)
	if err != nil {
		panic(err)
	}
	p2CardPoint, err := strconv.Atoi(covertP2Card)
	if err != nil {
		panic(err)
	}
	return p1CardPoint > p2CardPoint
}
