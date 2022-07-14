package meennew

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1WinWithTen(t *testing.T) {
	round := Round{value: "5S 9S TS 3H 2S 4H 2H 5H 6S 7S"}
	assert.Equal(t, round.isPlayer1Winner(), true)
}
