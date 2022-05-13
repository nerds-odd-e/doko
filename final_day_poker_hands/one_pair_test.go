package final_day_poker_hands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeP1WinOnePair() []string {
	return []string{"2S 3C 4D 8S 8H 2S 3D 4C 5H 9H"}
}

func makeP1WinOnePairAnotherFace() []string {
	return []string{"2S 3C 4D 6C 6D 2S 3D 4C 5H 9H"}
}

func TestP1WinsOnePair(t *testing.T) {
	assert.Equal(t, 1, pokerhands(makeP1WinOnePair()))
	assert.Equal(t, 1, pokerhands(makeP1WinOnePairAnotherFace()))
}
