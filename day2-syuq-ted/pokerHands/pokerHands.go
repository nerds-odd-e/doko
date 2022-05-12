package pokerhands

import "strings"

type Game []string

func pokerHands(hands []string) int {
	winCount := 0
	for i := range hands {
		if p1Wins(strings.Fields(hands[i])) {
			winCount += 1
		}
	}
	return winCount
}

func p1Wins(cards []string) bool {
	left, right := 4, 9
	for cards[left] == cards[right] && left > 0 && right > 5 {
		left -= 1
		right -= 1
	}
	if string(cards[right][0]) == "T" {
		return true
	}
	if string(cards[left][0]) == "T" || string(cards[left][0]) == "Q" {
		return false
	}

	return cards[left] > cards[right]
}
