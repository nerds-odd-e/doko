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
	records := []string{"5S AD TS 3H 2S 4H 2H 5H 6S QS"}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 1)
}

func TestTwoRecord(t *testing.T) {
	records := []string{
		"5S AD TS 3H 2S 4H 2H 5H 6S QS",
		"5S AD TS 3H 2S 4H 2H 5H 6S QS",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 2)
}

func TestPlayerOneLose(t *testing.T) {
	records := []string{
		"5S 6D TS 3H 2S 4H 2H 5H 6S QS",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 0)
}

func TestPlayerOneWinInTwoRound(t *testing.T) {
	records := []string{
		"5S 6D TS 3H 2S 4H 2H 5H 6S QS",
		"5S AD TS 3H 2S 4H 2H 5H 6S QS",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 0)
}
