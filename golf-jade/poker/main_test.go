package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHighCardWin(t *testing.T) {
	got := FindPokerWinPercentage([]string{"8C TS KC 9H 4S 7D 2S 5D 3S AC"})
	want := 100.0
	assert.Equal(t, got, want)
}
