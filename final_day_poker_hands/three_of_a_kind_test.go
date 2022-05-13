package final_day_poker_hands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeP1WinThreeOfAKind() []string {
	return []string{"2S 3C 8D 8S 8H 2S 3D 4C 5H 9H"}
}

func TestReturn1IfP1HasThreeOfAKind(t *testing.T) {
	assert.Equal(t, 1, pokerhands(makeP1WinThreeOfAKind()))
}
