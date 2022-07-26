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

func TestOneGameLoseWithHighCard(t *testing.T) {
	result := getFirstPlayerWinCount("./files/one_game_lose_high_card.txt")
	assert.Equal(t, result, 0, "Input file is one_game_lose_high_card.txt, should get 0")
}

func TestTwoGameWinWithHighCard(t *testing.T) {
	result := getFirstPlayerWinCount("./files/two_game_win_high_card.txt")
	assert.Equal(t, result, 2, "Input file is two_game_win_high_card.txt, should get 2")
}

func TestTwoGameLoseWithHighCard(t *testing.T) {
	result := getFirstPlayerWinCount("./files/two_game_lose_high_card.txt")
	assert.Equal(t, result, 0, "Input file is two_game_lose_high_card.txt, should get 0")
}

func TestOneGameWinWithHighCardIsNotFirstCard(t *testing.T) {
	result := isPlayer1Win("5C 4S KC 3H 2S QD 5S 4D 3S 2C")
	assert.Equal(t, result, false, "Input highest score card is in third index, should get true")
}
