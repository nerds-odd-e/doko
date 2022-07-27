package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWinrate(t *testing.T) {
	t.Run("Empty game", func(t *testing.T) {
		game := []string{}
		assert.Equal(t, 0.0, P1Winrate(game))
	})
	t.Run("P1 win one game", func(t *testing.T) {
		game := []string{"AC 9S 7D 3C 2C QS 5C 4H 3D 2C"}
		assert.Equal(t, 0.0, P1Winrate(game))
	})
}

func TestHighCard(t *testing.T) {

	t.Run("P1 win", func(t *testing.T) {
		game := "AC 6S 5D 3C 2C KS 5C 4H 3D 2C"
		assert.Equal(t, true, P1Win(game))
	})

	t.Run("P2 wins", func(t *testing.T) {
		game := "KS 5C 4H 3D 2C AC 6S 5D 3C 2C"
		assert.Equal(t, false, P1Win(game))
	})

	t.Run("P2 wins again.", func(t *testing.T) {
		game := "TS 5C 4H 3D 2C QC 6S 5D 3C 2C"
		assert.Equal(t, false, P1Win(game))
	})

	t.Run("P1 wins by J", func(t *testing.T) {
		game := "JS 5C 4H 3D 2C TC 6S 5D 3C 2C"
		assert.Equal(t, true, P1Win(game))
	})

	t.Run("P2 wins in other order", func(t *testing.T) {
		game := "4S 5C JH 3D 2C 2C 6S 5D QC 7C"
		assert.Equal(t, true, P1Win(game))
	})
}
func TestOpenEmptyFile_ReturnEmptyGame(t *testing.T) {
	assert.Equal(t, []string{}, OpenFile("data/empty_game.txt"))
}

func TestOpenOneLineFile_Return1Game(t *testing.T) {
	input := "data/one_line.txt"
	expected := []string{"AA AA AA AA AA AA AA AA AA AA"}

	assert.Equal(t, expected, OpenFile(input))
}

func TestOpenOneLineFile_ReturnAnother1Game(t *testing.T) {
	input := "data/one_line2.txt"
	expected := []string{"8C TS KC 9H 4S 7D 2S 5D 3S AC"}

	assert.Equal(t, expected, OpenFile(input))
}

func TestOpenTwoLineFile_Return2Games(t *testing.T) {
	input := "data/one_line2.txt"
	expected := []string{
		"8C TS KC 9H 4S 7D 2S 5D 3S AC",
	}

	assert.Equal(t, expected, OpenFile(input))
}
