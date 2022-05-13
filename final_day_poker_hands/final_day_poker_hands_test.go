package final_day_poker_hands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturn0WinsIfInputIsEmpty(t *testing.T) {
	assert.Equal(t, 0, pokerhands([]string{}))
}

func makeP1HighestCardHighCard() string {
	return "2S 3C 4D 5H 9H 6D 2S 3D 4C 5H"
}

func xTestReturn1IfP1HasHighestCardAtPosition5(t *testing.T) {
	assert.Equal(t, 1, pokerhands([]string{"8C TS KC 7H 9S 4S 7D 2S 5D 3S"}))
}
