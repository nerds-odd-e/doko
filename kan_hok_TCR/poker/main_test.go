package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFirstPlayerWinCount(t *testing.T) {
	result := getFirstPlayerWinCount("./files/empty.txt")
	assert.Equal(t, result, 0, "Input is empty file which is no game, should get 0")
}

func TestOneGameWinWithHighCard(t *testing.T) {
	result := getFirstPlayerWinCount("./files/one_game_win_high_card.txt")
	assert.Equal(t, result, 1, "Input file is one_game_win_high_card.txt, should get 1")
}
