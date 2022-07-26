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

func Test_input_lose_with_pair(t *testing.T) {
	assert.Equal(t, 0.0, Poker("pair_win_game.txt"))
}

func Test_input_another_lose_with_pair(t *testing.T) {
	assert.Equal(t, 0.0, Poker("pair_win_game_1.txt"))
}

func Test_pair(t *testing.T) {
	assert.Equal(t, true, player_one_win(Hand{[]string{"3H", "7H", "6H", "6S", "5D"}}, Hand{[]string{"8D", "AD", "JD", "TD", "9S"}}))
}

func Test_three_of_a_kind(t *testing.T) {
	assert.Equal(t, true, player_one_win(Hand{[]string{"3H", "7H", "6H", "6S", "6D"}}, Hand{[]string{"8D", "AD", "JD", "TD", "8S"}}))
	assert.Equal(t, false, player_one_win(Hand{[]string{"3H", "7H", "6H", "6S", "6D"}}, Hand{[]string{"8D", "AD", "7H", "7S", "7D"}}))
}
