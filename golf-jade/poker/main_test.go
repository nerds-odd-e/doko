package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHighCardWinA(t *testing.T) {
	got := FindPokerWinPercentage([]string{"AC TS KC 9H 4S 7D 2S 5D 3S KC"})
	want := 100.0
	assert.Equal(t, got, want)
}

func TestHighCardLoseA(t *testing.T) {
	got := FindPokerWinPercentage([]string{"KC TS 4C 9H 4S 7D 2S AD 3S KC"})
	want := 0.0
	assert.Equal(t, got, want)
}

func TestHighCardDrawA(t *testing.T) {
	got := FindPokerWinPercentage([]string{"KC TS 4C 9H 4S 7D 2S AD 3S KC", "AC TS KC 9H 4S 7D 2S 5D 3S KC"})
	want := 50.0
	assert.Equal(t, got, want)
}
