package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerAWinTheGame(t *testing.T) {
	hands := "8C TS KC 9H 4S 7D 2S 5D 3S AC"

	playerAWin := PokerGamePlayerAWin(hands)
	assert.Equal(t, true, playerAWin)
}


/*
func TestPlayerALooseTheGame(t *testing.T) {
	hands := "5C AD 5D AC 9C 7C 5H 8D TD KS"

	playerAWin := PokerGamePlayerAWin(hands)
	assert.Equal(t, false, playerAWin)
}
*/
