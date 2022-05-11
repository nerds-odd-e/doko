package pokerhands

import "strings"

func pokerHands(hands []string) int {
	if len(hands) == 1 {

		cards := strings.Fields(hands[0])

		if cards[4] > cards[9] {
			return 1
		}
	}

	return 0
}
