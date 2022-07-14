package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func Test_pokerEmptyRow(t *testing.T) {
// 	pokerFile := []string{}
// 	a := calculatorPercentWinnerPoker(pokerFile)
// 	b := float64(0)
// 	assert.Equal(t, a, b)
// }

// func Test_pokerOneRow(t *testing.T) {
// 	pokerFile := []string{
// 		"TH JH QH KH AH    6H 5H 2H 3H 4H",
// 	}
// 	a := calculatorPercentWinnerPoker(pokerFile)
// 	b := float64(100)
// 	assert.Equal(t, a, b)
// }

// func Test_pokerTwoRow(t *testing.T) {
// 	pokerFile := []string{
// 		"TH JH QH KH AH    6H 5H 2H 3H 4H",
// 		"TH JH QH KH AH    6H 5H 2H 3H 4H",
// 	}
// 	a := calculatorPercentWinnerPoker(pokerFile)
// 	b := float64(100)
// 	assert.Equal(t, a, b)
// }

// func Test_pokerTwoRowFifttyPercentWinRate(t *testing.T) {
// 	pokerFile := []string{
// 		"TH JH QH KH AH    6H 5H 2H 3H 4H",
// 		"6H 5H 2H 3H 4H    TH JH QH KH AH",
// 	}
// 	a := calculatorPercentWinnerPoker(pokerFile)
// 	b := float64(50)
// 	assert.Equal(t, a, b)
// }

// func Test_pokerThreeRowSixtySixPercentWinRate(t *testing.T) {
// 	pokerFile := []string{
// 		"TH JH QH KH AH    6H 5H 2H 3H 4H",
// 		"TH JH QH KH AH    6H 5H 2H 3H 4H",
// 		"6H 5H 2H 3H 4H    TH JH QH KH AH",
// 	}
// 	a := calculatorPercentWinnerPoker(pokerFile)
// 	b := float64(66.66666666666667)
// 	assert.Equal(t, a, b)
// }

// func Test_pokerTwoRowZeroWinRate(t *testing.T) {
// 	pokerFile := []string{
// 		"6H 5H 2H 3H 4H    TH JH QH KH AH",
// 		"6H 5H 2H 3H 4H    TH JH QH KH AH",
// 	}
// 	a := calculatorPercentWinnerPoker(pokerFile)
// 	b := float64(0)
// 	assert.Equal(t, a, b)
// }

// func Test_pokerTwoRow100WinRate(t *testing.T) {
// 	pokerFile := []string{
// 		"TH JH QH KH AH    6H 5H 2H 3H 4H",
// 		"TH JH QH KH AH    6H 5H 2H 3H 4H",
// 	}
// 	a := calculatorPercentWinnerPoker(pokerFile)
// 	b := float64(100)
// 	assert.Equal(t, a, b)
// }

// func Test_pokerRoyalflush(t *testing.T) {
// 	a := calculatorPokerHand("TH JH QH KH AH")
// 	b := 10
// 	assert.Equal(t, a, b)
// }

// func Test_pokerStraightflush(t *testing.T) {
// 	a := calculatorPokerHand("6H 5H 2H 3H 4H")
// 	b := 9
// 	assert.Equal(t, a, b)
// }

// func Test_pokerFourOfKind(t *testing.T) {
// 	a := calculatorPokerHand("9H 9H 9H 9H 4H")
// 	b := 8
// 	assert.Equal(t, a, b)
// }

// func Test_pokerFullhouse(t *testing.T) {
// 	a := calculatorPokerHand("9H 9C 9D AS AH")
// 	b := 7
// 	assert.Equal(t, a, b)
// }

// func Test_pokerFlush(t *testing.T) {
// 	a := calculatorPokerHand("AD JD 7D 9D AD")
// 	b := 6
// 	assert.Equal(t, a, b)
// }
// func Test_pokerStraight(t *testing.T) {
// 	a := calculatorPokerHand("9H 8C 7D 6S 5H")
// 	b := 5
// 	assert.Equal(t, a, b)
// }
// func Test_pokerThreeofKind(t *testing.T) {
// 	a := calculatorPokerHand("5H QC 2D 2S 2H")
// 	b := 4
// 	assert.Equal(t, a, b)
// }
// func Test_pokerTwoPair(t *testing.T) {
// 	a := calculatorPokerHand("AH AC KD KS 9H")
// 	b := 3
// 	assert.Equal(t, a, b)
// }
// func Test_pokerOnePair(t *testing.T) {
// 	a := calculatorPokerHand("4H 5C 8D AS AH")
// 	b := 2
// 	assert.Equal(t, a, b)
// }
// func Test_pokerHighCard_1(t *testing.T) {
// 	a := calculatorPokerHand("5H 6C JD QS AH")
// 	b := 1
// 	assert.Equal(t, a, b)
// }

// func xTest_pokerHighCard_2(t *testing.T) {
// 	a := calculatorPokerHand("JH 6C JD QS QH")
// 	b := 1
// 	assert.Equal(t, a, b)
// }

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

func TestP1WinByHighcard(t *testing.T) {
	assert.Equal(t, calculatorPercentWinnerPoker(aGameOfTwoHighCards("A", "K")), 1)
	assert.Equal(t, calculatorPercentWinnerPoker(aGameOfTwoHighCards("K", "Q")), 1)
	assert.Equal(t, calculatorPercentWinnerPoker(aGameOfTwoHighCards("Q", "J")), 1)
	assert.Equal(t, calculatorPercentWinnerPoker(aGameOfTwoHighCards("J", "T")), 1)
	assert.Equal(t, calculatorPercentWinnerPoker(aGameOfTwoHighCards("T", "9")), 1)
}

func TestPlayer1LostByHighCard(t *testing.T) {
	assert.Equal(t, calculatorPercentWinnerPoker(aGameOfTwoHighCards("K", "A")), 0)
	assert.Equal(t, calculatorPercentWinnerPoker(aGameOfTwoHighCards("Q", "A")), 0)
	assert.Equal(t, calculatorPercentWinnerPoker(aGameOfTwoHighCards("J", "A")), 0)
	assert.Equal(t, calculatorPercentWinnerPoker(aGameOfTwoHighCards("T", "A")), 0)
}
