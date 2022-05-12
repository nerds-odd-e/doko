package pokerHands

import (
	"strings"
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

func Test_poker_p1_wins_high_card_2_games(t *testing.T) {
	p1 := "9C 2D 2C 3D 3C"
	p2 := "8S 2K 2S 3K 3S"
	games := p1 + " " + p2 + "\n" + p1 + " " + p2
	assert.Equal(t, 2, p1_win_counts(games))
}

func p1_win_counts(games string) int {

	wins := 0
	for _, game := range toGames(games) {
		if check_if_p1_win_high_card(game) {
			wins += 1
		}
	}
	return wins
}

func toGames(games string) []string {
	if games == "" {
		return []string{}
	}

	return strings.Split(games, "\n")
}

func check_if_p1_win_high_card(game string) bool {
	return get_high_card(game, "p1") > get_high_card(game, "p2")
}

func get_high_card(games string, player string) byte {
	if player == "p1" {
		return games[0]
	}
	return games[15]
}
