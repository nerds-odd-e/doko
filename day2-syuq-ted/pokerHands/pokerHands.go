package pokerhands

import "strings"

type Game []string

func pokerHands(hands []string) int {
	winCount := 0
	for i := range hands {
		game := Game(strings.Fields(hands[i]))
		if game.p1Wins() {
			winCount += 1
		}
	}
	return winCount
}
func (game Game) p1Wins() bool {
	left, right := 4, 9
	for game[left] == game[right] && left > 0 && right > 5 {
		left -= 1
		right -= 1
	}
	if string(game[right][0]) == "T" {
		return true
	}
	if string(game[left][0]) == "T" || string(game[left][0]) == "Q" {
		return false
	}

	return game[left] > game[right]
}
