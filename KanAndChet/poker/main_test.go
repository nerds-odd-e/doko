package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyFile(t *testing.T) {
	result := getFirstPlayerWinRate("empty.txt")
	assert.Equal(t, result, 0.0, "The empty file should return 0.0")
}

func TestOneGameFileWithFirstPlayerWin(t *testing.T) {
	result := getFirstPlayerWinRate("one_game_with_win.txt")
	assert.Equal(t, result, 1.0, "In one_game_with_win, the first player win, so the win rate is 1.0")
}

func TestOneGameWithFirstplayerLose(t *testing.T) {
	result := getFirstPlayerWinRate("one_game_with_lose.txt")
	assert.Equal(t, result, 0.0, "In one_game_with_lose.txt, the first player lose, so the win rate is 0.0")
}

func TestNGameWithFirstPlayerWinLose(t *testing.T) {
	result := getFirstPlayerWinRate("two_game_win_lose.txt")
	assert.Equal(t, result, 0.5, "In two_game_win_lose.txt, the first player win rate should be 0.5")
}
