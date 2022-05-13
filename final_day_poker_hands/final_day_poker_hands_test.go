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

func xTestReturn1IfP1HasHighestCardIfHandIsSorted(t *testing.T){
	assert.Equal(t,1,pokerhands([]string{makeP1HighestCardHighCard()}))
}

func xTestReturn0IfP2HasHighestCardIfHandIsSorted(t *testing.T){
	assert.Equal(t,0,pokerhands([]string{"2C 3H 4C 5C 6C 2S 3S 4S 5S AS"}))
}
