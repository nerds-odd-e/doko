package pokerhand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Input_empty_File(t *testing.T) {
	assert.Equal(t, 0.0, Poker("empty.txt"))
}

func Test_Input_one_game_file(t *testing.T) {
	assert.Equal(t, 1.0, Poker("one_game.txt"))
}

func Test_Input_one_game_another_file(t *testing.T) {
	assert.Equal(t, 1.0, Poker("one_game_2.txt"))
}

func Test_input_one_games_lose(t *testing.T) {
	assert.Equal(t, 0.0, Poker("one_game_3.txt"))
}

func Test_input_another_one_games_lose(t *testing.T) {
	assert.Equal(t, 0.0, Poker("one_game_4.txt"))
}

func Test_input_multiple_games(t *testing.T) {
	assert.Equal(t, 0.5, Poker("multiple_game.txt"))
}

func Test_input_multiple_games_all_win(t *testing.T) {
	assert.Equal(t, 1.0, Poker("multiple_game_1.txt"))
}
