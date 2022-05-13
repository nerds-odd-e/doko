package final_day_poker_hands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeP1WinHighestCardHighCardTwoGames() []string {
	return []string{"2S 3C 4D 5H 9H 2S 3D 4C 5H 6D", "2S 3C 4D 5H 9H 2S 3D 4C 5H 6D"}
}

func TestP1WinsHighCardTwiceOfTwoGames(t *testing.T) {
	assert.Equal(t, 2, pokerhands(makeP1WinHighestCardHighCardTwoGames()))
}