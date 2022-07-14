package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerWin0Game(t *testing.T) {
	a := []string{}
	assert.Equal(t, calculatorPercentWinnerPoker(a), 0)
}

func TestPlayerWin1Game(t *testing.T) {
	a := []string{
		"9H 9C 9D AS 9H    8H 9C 9D AS 8H",
	}
	assert.Equal(t, calculatorPercentWinnerPoker(a), 1)
}
func TestPlayerWin2Game(t *testing.T) {
	a := []string{
		"9H 9C 9D AS 9H    8H 9C 9D AS 8H",
		"9H 9C 9D AS 9H    8H 9C 9D AS 8H",
	}
	assert.Equal(t, calculatorPercentWinnerPoker(a), 2)
}

func TestPlayerWin1in2Game(t *testing.T) {
	a := []string{
		"8H 9C 9D AS 9H    7H 9C 9D AS 9H",
		"8H 9C 9D AS 9H    9H 9C 9D AS 9H",
	}
	assert.Equal(t, calculatorPercentWinnerPoker(a), 1)
}

func TestPlayerWin2in2Game(t *testing.T) {
	a := []string{
		"9H 9C 9D AS 9H    8H 9C 9D AS 9H",
		"9H 9C 9D AS 9H    8H 9C 9D AS 9H",
	}
	assert.Equal(t, calculatorPercentWinnerPoker(a), 2)
}

type PokerHandBuilder struct {
	highestHighcard string
}

func pokerHand() PokerHandBuilder {
	return PokerHandBuilder{}
}

func (b PokerHandBuilder) highCardOf(card string) PokerHandBuilder {
	b.highestHighcard = card
	return b
}

func (b PokerHandBuilder) please() string {
	return b.highestHighcard + "H 2C 3D 4S 6H"
}

func aGameOfTwoHighCards(p1HighCard string, p2HighCard string) []string {
	return []string{pokerHand().highCardOf(p1HighCard).please() + "    " + pokerHand().highCardOf(p2HighCard).please()}
}

func twoGameOfTwoHighCards(p1HighCard []string, p2HighCard []string) []string {
	return []string{
		pokerHand().highCardOf(p1HighCard[0]).please() + "    " + pokerHand().highCardOf(p2HighCard[0]).please(),
		pokerHand().highCardOf(p1HighCard[1]).please() + "    " + pokerHand().highCardOf(p2HighCard[1]).please(),
	}
}

func TestP1WinByHighcard(t *testing.T) {
	assert.Equal(t, calculatorPercentWinnerPoker(aGameOfTwoHighCards("A", "K")), 1)
	assert.Equal(t, calculatorPercentWinnerPoker(aGameOfTwoHighCards("K", "Q")), 1)
	assert.Equal(t, calculatorPercentWinnerPoker(aGameOfTwoHighCards("Q", "J")), 1)
	assert.Equal(t, calculatorPercentWinnerPoker(aGameOfTwoHighCards("J", "T")), 1)
	assert.Equal(t, calculatorPercentWinnerPoker(aGameOfTwoHighCards("T", "9")), 1)
	assert.Equal(t, calculatorPercentWinnerPoker(twoGameOfTwoHighCards([]string{"T", "J"}, []string{"9", "T"})), 2)
}

func TestPlayer1LostByHighCard(t *testing.T) {
	assert.Equal(t, calculatorPercentWinnerPoker(aGameOfTwoHighCards("K", "A")), 0)
	assert.Equal(t, calculatorPercentWinnerPoker(aGameOfTwoHighCards("Q", "A")), 0)
	assert.Equal(t, calculatorPercentWinnerPoker(aGameOfTwoHighCards("J", "A")), 0)
	assert.Equal(t, calculatorPercentWinnerPoker(aGameOfTwoHighCards("T", "A")), 0)
	assert.Equal(t, calculatorPercentWinnerPoker(twoGameOfTwoHighCards([]string{"T", "J"}, []string{"A", "A"})), 0)
}
