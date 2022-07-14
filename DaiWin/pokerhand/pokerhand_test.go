package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const highCardWithEigth = "2H 3H 4H 5H 8S"
const highCardWithNine = "2H 3H 4H 5H 9S"
const highCardWithKing = "2H 3H 4H 5H KS"
const highCardWithAce = "2H 3H 5H 4H AS"
const highCardWithQueen = "2H 3H 4H 5H QS"
const highCardWithJack = "2H 3H 4H JS 5H"
const highCardWithTen = "2H TS 3H 4H 5H"
const highCardWithSeven = "2H 3H 4H 5H 7S"

func TestPlayer1Win0TimesWhenNoGamePlayed(t *testing.T) {
	pokerFile := []string{}
	assert.Equal(t, 0, CalculatePlayer1WinGame(pokerFile))
}

func TestPlayer1WinOnceIn1Game(t *testing.T) {
	pokerFile := []string{highCardWithEigth + "  " + highCardWithSeven}
	assert.Equal(t, 1, CalculatePlayer1WinGame(pokerFile))
}
func TestPlayer1WinOnceIn1Game1(t *testing.T) {
	pokerFile := []string{highCardWithNine + "  " + highCardWithSeven}
	assert.Equal(t, 1, CalculatePlayer1WinGame(pokerFile))
}

func TestPlayer1LoseOnceIn1Game(t *testing.T) {
	pokerFile := []string{highCardWithSeven + "  " + highCardWithEigth}
	assert.Equal(t, 0, CalculatePlayer1WinGame(pokerFile))
}

func TestPlayer1Win2In2Game(t *testing.T) {
	pokerFile := []string{highCardWithEigth + "  " + highCardWithSeven, highCardWithEigth + "  " + highCardWithSeven}
	assert.Equal(t, 2, CalculatePlayer1WinGame(pokerFile))
}

func TestPlayer1WinOnceIn2Game(t *testing.T) {
	pokerFile := []string{highCardWithEigth + "  " + highCardWithSeven, highCardWithSeven + "  " + highCardWithEigth}
	assert.Equal(t, 1, CalculatePlayer1WinGame(pokerFile))
}

func TestPlayer1LoseOnceIn2Game(t *testing.T) {
	pokerFile := []string{highCardWithSeven + "  " + highCardWithEigth, highCardWithEigth + "  " + highCardWithSeven}
	assert.Equal(t, 1, CalculatePlayer1WinGame(pokerFile))
}

func TestPlayer1WinOnceUn1GameWithKingCard(t *testing.T) {
	pokerFile := []string{highCardWithKing + "  " + highCardWithSeven}
	assert.Equal(t, 1, CalculatePlayer1WinGame(pokerFile))
}

func TestPlayer1WinOnceUn1GameWithAceCard(t *testing.T) {
	pokerFile := []string{highCardWithAce + "  " + highCardWithKing}
	assert.Equal(t, 1, CalculatePlayer1WinGame(pokerFile))
}

func TestPlayer1WinOnceUn1GameWithQueenCard(t *testing.T) {
	pokerFile := []string{highCardWithQueen + "  " + highCardWithJack}
	assert.Equal(t, 1, CalculatePlayer1WinGame(pokerFile))
}

func TestPlayer1WinOnceUn1GameWithJackCardAndUnSortHighCard(t *testing.T) {
	pokerFile := []string{highCardWithJack + "  " + highCardWithTen}
	assert.Equal(t, 1, CalculatePlayer1WinGame(pokerFile))
}

func TestPlayer1WinOnceUn1GameWithTenCardAndUnSortHighCard(t *testing.T) {
	pokerFile := []string{highCardWithTen + "  " + highCardWithNine}
	assert.Equal(t, 1, CalculatePlayer1WinGame(pokerFile))
}

type Game struct {
	game string
}

func CalculatePlayer1WinGame(list []string) int {
	player1WinningCount := 0
	for _, row := range list {
		if isPlayerOneWin(Game{row}) {
			player1WinningCount += 1
		}
	}
	return player1WinningCount
}

func isPlayerOneWin(game Game) bool {
	cardInHand := strings.Split(game.game, "  ")
	return findHighCardPointInHand(cardInHand[0]) > findHighCardPointInHand(cardInHand[1])
}
