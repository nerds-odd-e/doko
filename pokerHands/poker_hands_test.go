package pokerHands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_poker_empty(t *testing.T) {
	game := ""
	assert.Equal(t, 0, p1_win_counts(game))
}

func Test_poker_p1_wins_high_card(t *testing.T) {
	p1 := "9C 2D 2C 3D 3C"
	p2 := "8S 2K 2S 3K 3S"
	game := p1 + " " + p2
	assert.Equal(t, 1, p1_win_counts(game))
}

func Test_poker_p2_wins_high_card(t *testing.T) {
	p1 := "8C 2D 2C 3D 3C"
	p2 := "9S 2K 2S 3K 3S"
	game := p1 + " " + p2
	assert.Equal(t, 0, p1_win_counts(game))
}

// func Test_poker_p1_wins_high_card_2_games(t *testing.T) {
// 	p1 := "9C 2D 2C 3D 3C"
// 	p2 := "8S 2K 2S 3K 3S"
// 	p1_2:= "9C 2D 2C 3D 3C"
// 	p2_2 := "8S 2K 2S 3K 3S"
// 	hands := p1 + " " + p2 + "\n"+ p1 + " " + p2
// 	assert.Equal(t, 2, p1_win_counts(hands))
// }

func p1_win_counts(games string) int {

	if games == "" {
		return 0
	}

	// game_list := strings.Split(games, "\n")

	// for game :=

	if get_high_card(games, "p1") > get_high_card(games, "p2") {
		return 1
	}
	return 0
}

func get_high_card(games string, player string) byte {
	if player == "p1" {
		return games[0]
	}
	return games[15]
}
