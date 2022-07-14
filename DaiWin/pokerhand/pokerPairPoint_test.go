package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const threePointPairCard = "3H 3H TS 4H 5H"
const twoPointPairCard = "2H 2H TS 4H 5H"

func TestPairCardInHand(t *testing.T) {
	pokerFile := []string{threePointPairCard + "  " + twoPointPairCard}
	assert.Equal(t, 0, CalculatePlayer1WinGame(pokerFile))
}
