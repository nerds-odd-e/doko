package pokerHands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_poker_empty(t *testing.T) {
	hands := ""
	assert.Equal(t, 0, poker(hands))
}

func poker(hands string) int {
	return 0
}
