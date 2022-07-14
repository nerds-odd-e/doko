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

const highCardWithEigth = "2H 3H 4H 5H 8S"
const highCardWithNine = "2H 3H 4H 5H 9S"
const highCardWithKing = "2H 3H 4H 5H KS"
const highCardWithAce = "2H 3H 4H 5H AS"
const highCardWithQueen = "2H 3H 4H 5H QS"
const highCardWithJack = "2H 3H 4H 5H JS"
const highCardWithTen = "2H 3H 4H 5H TS"
const highCardWithSeven = "2H 3H 4H 5H 7S"

func TestPlayer1Win0TimesWhenNoGamePlayed(t *testing.T) {
	pokerFile := []string{}
	assert.Equal(t, 0, ValidatePokerFile(pokerFile))
}

func TestPlayer1WinOnceIn1Game(t *testing.T) {
	pokerFile := []string{highCardWithEigth + " " + highCardWithSeven}
	assert.Equal(t, 1, ValidatePokerFile(pokerFile))
}
func TestPlayer1WinOnceIn1Game1(t *testing.T) {
	pokerFile := []string{highCardWithNine + " " + highCardWithSeven}
	assert.Equal(t, 1, ValidatePokerFile(pokerFile))
}

func TestPlayer1WinOnceIn2Game(t *testing.T) {
	pokerFile := []string{highCardWithEigth + " " + highCardWithSeven, highCardWithSeven + " " + highCardWithEigth}
	assert.Equal(t, 1, ValidatePokerFile(pokerFile))
}

func TestPlayer1LoseOnceIn1Game(t *testing.T) {
	pokerFile := []string{highCardWithSeven + " " + highCardWithEigth}
	assert.Equal(t, 0, ValidatePokerFile(pokerFile))
}

func TestPlayer1LoseOnceIn2Game(t *testing.T) {
	pokerFile := []string{highCardWithSeven + " " + highCardWithEigth, highCardWithEigth + " " + highCardWithSeven}
	assert.Equal(t, 1, ValidatePokerFile(pokerFile))
}

func TestPlayer1Win2In2Game(t *testing.T) {
	pokerFile := []string{highCardWithEigth + " " + highCardWithSeven, highCardWithEigth + " " + highCardWithSeven}
	assert.Equal(t, 2, ValidatePokerFile(pokerFile))
}

func TestPlayer1Win2In3Game(t *testing.T) {
	pokerFile := []string{highCardWithEigth + " " + highCardWithSeven, highCardWithEigth + " " + highCardWithSeven, highCardWithSeven + " " + highCardWithSeven}
	assert.Equal(t, 2, ValidatePokerFile(pokerFile))
}

func TestPlayer1WinOnceUn1GameWithKingCard(t *testing.T) {
	pokerFile := []string{highCardWithKing + " " + highCardWithSeven}
	assert.Equal(t, 1, ValidatePokerFile(pokerFile))
}

func TestPlayer1WinOnceUn1GameWithAceCard(t *testing.T) {
	pokerFile := []string{highCardWithAce + " " + highCardWithKing}
	assert.Equal(t, 1, ValidatePokerFile(pokerFile))
}

func TestPlayer1WinOnceUn1GameWithQueenCard(t *testing.T) {
	pokerFile := []string{highCardWithQueen + " " + highCardWithJack}
	assert.Equal(t, 1, ValidatePokerFile(pokerFile))
}

func TestPlayer1WinOnceUn1GameWithJackCard(t *testing.T) {
	pokerFile := []string{highCardWithJack + " " + highCardWithTen}
	assert.Equal(t, 1, ValidatePokerFile(pokerFile))
}

func TestPlayer1WinOnceUn1GameWithTenCard(t *testing.T) {
	pokerFile := []string{highCardWithTen + " " + highCardWithNine}
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
	mapPointWithHonorCard := map[string]string{
		"A": "14",
		"K": "13",
		"Q": "12",
		"J": "11",
		"T": "10",
		"9": "9",
		"8": "8",
		"7": "7",
	}
	covertP1Card := string(row[12])
	covertP2Card := string(row[27])
	covertP1Card = mapPointWithHonorCard[covertP1Card]
	covertP2Card = mapPointWithHonorCard[covertP2Card]

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
