package pokerhands

import "strings"

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
	return cards[left] > cards[right]
}
