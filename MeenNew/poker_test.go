package meennew

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
		"5S 6D TS 3H 2S 4H 2H 5H 6S 9S",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 0)
}

func TestPlayerOneWinInTwoRound(t *testing.T) {
	records := []string{
		"5S 6D TS 3H 2S 4H 2H 5H 6S 9S",
		"5S AD TS 3H 2S 4H 2H 5H 6S 9S",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 1)
}

func TestPlayer2WinOf3Round(t *testing.T) {
	records := []string{
		"5S 6D TS 3H 2S 4H 2H 5H 6S 9S",
		"5S AD TS 3H 2S 4H 2H 5H 6S 9S",
		"5S AD TS 3H 2S 4H 2H 5H 6S 9S",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 2)
}

func TestCompareP1WinHighCard(t *testing.T) {
	records := []string{
		"5S 6D TS 3H 2S 4H 2H 5H 6S 9S",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 0)
}
