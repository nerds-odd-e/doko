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
	records := []string{"5S TD TS 3H 2S 4H 2H AH 6S QS"}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 1)
}

func TestTwoRecord(t *testing.T) {
	records := []string{
		"5S TD TS 3H 2S 4H 2H AH 6S QS",
		"5S TD TS 3H 2S 4H 2H AH 6S QS",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 2)
}

func TestPlayerOneWin1Time(t *testing.T) {
	records := []string{
		"5S 6D 7S 8H TS 4H 2H AH 6S 3S",
		"5S 6D 7S 8H 3S 4H 2H AH 6S 9S",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 1)
}
