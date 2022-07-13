package meennew

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoRecord(t *testing.T) {
	records := []string{}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 0.0)
}

func TestOneRecord(t *testing.T) {
	records := []string{"5S TD TS 3H 2S", "4H 2H AH 6S QS"}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 100.0)
}

func TestTwoRecord(t *testing.T) {
	records := []string{"5S TD TS 3H 2S", "4H 2H AH 6S QS", "5S TD TS 3H 2S", "4H 2H AH 6S QS"}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 100.0)
}

func TestThreeRecord(t *testing.T) {
	records := []string{"5S TD TS 3H 2S", "4H 2H AH 6S QS", "5S TD TS 3H 2S", "4H 2H AH 6S QS", "4H 2H AH 6S QS", "5S TD TS 3H 2S"}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 66.66)
}

func TestThreeRecordP1OneWin(t *testing.T) {
	records := []string{"5S TD TS 3H 2S", "4H 2H AH 6S QS", "4H 2H AH 6S QS", "5S TD TS 3H 2S", "4H 2H AH 6S QS", "5S TD TS 3H 2S"}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 33.33)
}
