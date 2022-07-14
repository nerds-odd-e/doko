package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PokerHandNoGamePlay(t *testing.T) {
	pokerFile := []string{}
	a := calculatorWinGamePokerHand(pokerFile)
	b := 0
	assert.Equal(t, a, b)
}

func calculatorWinGamePokerHand([]string) int {
	return 0
}
