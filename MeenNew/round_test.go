package meennew

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1WinWithTen(t *testing.T) {
	round := Round{value: "5S 9S TS 3H 2S 4H 2H 5H 6S 7S"}
	assert.Equal(t, round.isPlayer1Winner(), true)
}

func TestP1WinWithKing(t *testing.T) {
	round := Round{value: "5S 9S KS 3H 2S 4H 2H 5H 6S 7S"}
	assert.Equal(t, round.isPlayer1Winner(), true)
}

func TestP1WinWithQueen(t *testing.T) {
	round := Round{value: "5S 9S QS 3H 2S 4H 2H 5H 6S 7S"}
	assert.Equal(t, round.isPlayer1Winner(), true)
}

func TestP1WinWithJack(t *testing.T) {
	round := Round{value: "5S 9S JS 3H 2S 4H 2H 5H 6S 7S"}
	assert.Equal(t, round.isPlayer1Winner(), true)
}

func TestP1WinWithNine(t *testing.T) {
	round := Round{value: "5S 9S 8S 3H 2S 4H 2H 5H 6S 7S"}
	assert.Equal(t, round.isPlayer1Winner(), true)
}

func TestP1WinWithEight(t *testing.T) {
	round := Round{value: "5S 8S 7S 3H 2S 4H 2H 5H 6S 7S"}
	assert.Equal(t, round.isPlayer1Winner(), true)
}

func TestP1WinWithSeven(t *testing.T) {
	round := Round{value: "7S 6S 5S 3H 2S 4H 2H 5H 6S 1S"}
	assert.Equal(t, round.isPlayer1Winner(), true)
}

func TestP1DrawWithAJ(t *testing.T) {
	round := Round{value: "AC JC 5S 3H 2S 4C 2H 5H AH JH"}
	assert.Equal(t, round.isPlayer1Winner(), false)
}

func TestP1DrawWithTen(t *testing.T) {
	round := Round{value: "TC 9C 5S 3H 2S 4C 2H 5H TH 9H"}
	assert.Equal(t, round.isPlayer1Winner(), false)
}

func TestP1Twopair(t *testing.T) {
	round := Round{value: "2C TD 5S 3H 2S 4C 2H 5H TH 9H"}
	assert.Equal(t, round.isPlayer1Winner(), false)
}
