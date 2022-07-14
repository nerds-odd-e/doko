package models

import (
	"sort"
	"strconv"
	"strings"
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
const unSortCard = "2H 3H TS 4H 5H"

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

func TestSortCardInHand(t *testing.T) {
	assert.Equal(t, "TS", sortCardInHand(unSortCard))
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
		"6": "6",
		"5": "5",
		"4": "4",
		"3": "3",
		"2": "2",
	}
	p1CardPoint, _ := strconv.Atoi(mapPointWithHonorCard[string(row[12])])
	p2CardPoint, _ := strconv.Atoi(mapPointWithHonorCard[string(row[27])])
	return p1CardPoint > p2CardPoint
}

func sortCardInHand(hand string) string {
	//hand = "2H 3H TS 4H 5H"
	splitStr := strings.Split(hand, " ")
	sort.Strings(splitStr)
	return splitStr[4]
}
