package meennew

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatasetNotFound(t *testing.T) {
	records := []string{}
	winRateP1 := SummaryWinner(records)
	assert.Equal(t, winRateP1, 0)
}

func TestPlayer1(t *testing.T) {
	records := []string{}
	winRateP1 := SummaryWinner(records)
	assert.Equal(t, winRateP1, 0)
}

func TestNoRecord(t *testing.T) {
	records := []string{}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 0)
}

func TestOneRecord(t *testing.T) {
	records := []string{"5S TD TS 3H 2S", "4H 2H AH 6S QS"}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 1)
}
