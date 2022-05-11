package pokerHands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_poker_empty(t *testing.T) {
	hands := ""
	assert.Equal(t, 0, p1_win_counts(hands))
}

func Test_poker_p1_wins_high_card(t *testing.T) {
	hands := "9C 2S 2C 2H 4S 2D 8S 3D 3S 3C"
	assert.Equal(t, 1, p1_win_counts(hands))
}

func p1_win_counts(hands string) int {
	if hands == "" {
		return 0
	}
	return 1
}
