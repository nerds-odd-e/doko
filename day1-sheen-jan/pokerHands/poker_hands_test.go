package pokerHands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoGames(t *testing.T) {
	assert.Equal(t, 0, pokerHands([]string{}))
}

func TestHighCard(t *testing.T) {
	//
	assert.Equal(t, 0, pokerHands([]string{"2C 3D 4D 5C 6C 2C 3H 4D 5D 7S"}))
	assert.Equal(t, 1, pokerHands([]string{"2C 3D 4D 5C 7C 2C 3H 4D 5D 6S"}))

	// One Game Unsorted
	assert.Equal(t, 0, pokerHands([]string{"7C 3D 4D 5C 2C 2C 6H 4D 5D 9S"}))
	assert.Equal(t, 1, pokerHands([]string{"9C 3D 4D 5C 2C 2C 6H 4D 5D 7S"}))
	// OneGameSameRank
	assert.Equal(t, 0, pokerHands([]string{"9C 3D 4D 5C 2C 2C 6H 4D 5D 9S"}))
	assert.Equal(t, 1, pokerHands([]string{"9S 3D 4D 5C 2C 2C 6H 4D 5D 9C"}))
	// OneGameWithCourtCards
	assert.Equal(t, 0, pokerHands([]string{"9C JD 4D 5C 2C KC 6H 4D 5D 8S"}))
	assert.Equal(t, 1, pokerHands([]string{"3S AD 4D 5C 2C 2C 6H QD 5D 9C"}))

	// OneGamme
	assert.Equal(t, 0, pokerHands([]string{"2C 3D 4D 5C 6C 2C 3H 4D 5D 7S"}))
	assert.Equal(t, 1, pokerHands([]string{"2C 3D 4D 5C 7C 2C 3H 4D 5D 6S"}))

	// Two Games
	assert.Equal(t, 0, pokerHands([]string{"9C JD 4D 5C 2C KC 6H 4D 5D 8S", "9C JD 4D 5C 2C KC 6H 4D 5D 8S"}))
	assert.Equal(t, 1, pokerHands([]string{"3S AD 4D 5C 2C 2C 6H QD 5D 9C", "9C JD 4D 5C 2C KC 6H 4D 5D 8S"}))
	assert.Equal(t, 2, pokerHands([]string{"3S AD 4D 5C 2C 2C 6H QD 5D 9C", "3S AD 4D 5C 2C 2C 6H QD 5D 9C"}))
}
