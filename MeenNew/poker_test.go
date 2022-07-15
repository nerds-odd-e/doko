package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoRecord(t *testing.T) {
	records := []string{}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 0)
}

func TestOneRecord(t *testing.T) {
	records := []string{"5S AD TS 3H 2S 4H 2H 5H 6S 9S"}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 1)
}

func TestTwoRecord(t *testing.T) {
	records := []string{
		"5S AD TS 3H 2S 4H 2H 5H 6S 9S",
		"5S AD TS 3H 2S 4H 2H 5H 6S 9S",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 2)
}

func TestPlayerOneLose(t *testing.T) {
	records := []string{
		"5S 6D TS 3H 2S 4H 2H 5H 6S KS",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 0)
}

func TestPlayerOneWinInTwoRound(t *testing.T) {
	records := []string{
		"5S 6D TS 3H 2S 4H 2H 5H 6S QS",
		"5S AD TS 3H 2S 4H 2H 5H 6S 9S",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 1)
}

func TestPlayer2WinOf3Round(t *testing.T) {
	records := []string{
		"5S 6D TS 3H 2S 4H 2H 5H 6S QS",
		"5S AD TS 3H 2S 4H 2H 5H 6S 9S",
		"5S AD TS 3H 2S 4H 2H 5H 6S 9S",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 2)
}

func TestCompareP1WinHighCard(t *testing.T) {
	records := []string{
		"5S KD TS 3H 2S 4H 2H 5H 6S 9S",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 1)
}

func TestCompareP1WinHighCardWithQueen(t *testing.T) {
	records := []string{
		"5S QD TS 3H 2S 4H 2H 5H 6S 9S",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 1)
}

func TestCompareP2WinHighCardWithKing(t *testing.T) {
	records := []string{
		"5S 4S TS 3H 2S 4H 2H QH 6S AS",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 0)
}

func TestCompareP1WinHighCardWithJack(t *testing.T) {
	records := []string{
		"5S JS TS 3H 2S 4H 2H 5H 6S TS",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 1)
}

func TestCompareP1WinHighCardWithTen(t *testing.T) {
	records := []string{
		"5S 4S TS 3H 2S 4H 2H 5H 6S 8S",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 1)
}

func TestCompareP1WinHighCardWithSix(t *testing.T) {
	records := []string{
		"5S 4S 3S 6H 2S 4H 2H 5H 2S 3S",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 1)
}

func TestCompareP1Draw(t *testing.T) {
	records := []string{
		"5S KD TS 3H 2S 4S KD 6S 3H 2S",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 0)
}
