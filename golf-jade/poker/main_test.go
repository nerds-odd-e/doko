package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const p2WinByHighCard = "3C TS 4C 9H 6S 7D 2S AD 3S 8C"
const p1WinByHighCard = "3H 7H 6S KC JS 5H TD JC 2D 8S"
const drawGameCard = "3H 7H 6S KC JS JH KD 7C 6D 3S"
const p1WinByHighNumberCard = "2H 4H 6S 8C 9S 8H 7D 6C 5D 4S"
const p1WinByHighCardSecondCard = "2H 4H 6S TC AS 8H AD 6C JD 4S"

func TestHighCardWinA(t *testing.T) {
	got := FindPokerWinPercentage([]string{"AC TS 8C 9H 6S 7D 2S 5D 3S 4C"})
	want := 100.0
	assert.Equal(t, want, got)
}

func TestHighCardLoseA(t *testing.T) {
	got := FindPokerWinPercentage([]string{p2WinByHighCard})
	want := 0.0
	assert.Equal(t, want, got)
}

func TestHighCardDrawA(t *testing.T) {
	got := FindPokerWinPercentage([]string{"8C TS 4C 9H 6S 7D 2S AD 3S 8D", "AC TS KC 9H 4S 7D 2S 5D 3S 8D"})
	want := 50.0
	assert.Equal(t, want, got)
}

func TestHighCardWinK(t *testing.T) {
	got := FindPokerWinPercentage([]string{"KC TS 4C 9H 2S 7D 2S 4D 3S 5C", "QC TS KC 9H 4S 7D 2S 5D 3S 7C"})
	want := 100.0
	assert.Equal(t, want, got)
}

func TestHighCardLoseK(t *testing.T) {
	got := FindPokerWinPercentage([]string{"KC TS 4C 9H 2S 7D 2S 4D 3S 5C", "QC TS KC 9H 4S 7D 2S 5D 3S 7C"})
	want := 100.0
	assert.Equal(t, want, got)
}

func TestHighCardWinQ(t *testing.T) {
	got := FindPokerWinPercentage([]string{"QC TS 4C 9H 2S 7D 2S 4D 6S 5C"})
	want := 100.0
	assert.Equal(t, want, got)
}

func TestHighCardLoseA2(t *testing.T) {
	got := FindPokerWinPercentage([]string{p1WinByHighCard, p2WinByHighCard, p2WinByHighCard})
	want := 33.33
	assert.Equal(t, want, got)
}

func TestDrawGame(t *testing.T) {
	got := FindPokerWinPercentage([]string{drawGameCard})
	want := 50.0
	assert.Equal(t, want, got)
}

func TestWinByHighNumberCard(t *testing.T) {
	got := FindPokerWinPercentage([]string{p1WinByHighNumberCard})
	want := 100.0
	assert.Equal(t, want, got)
}

func xTestLoseByHighCardSecondCard(t *testing.T) {
	got := FindPokerWinPercentage([]string{p1WinByHighCardSecondCard})
	want := 0.0
	assert.Equal(t, want, got)
}
