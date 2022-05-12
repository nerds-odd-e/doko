package exercises_qe_jt

import "strings"

const (
	HIGH_CARD_HIGHEST                   string = "3S 4H 5S 8S 9S"
	HIGH_CARD_SAME_HIGHEST_LOWER_SECOND string = "3S 4H 5S 7S 9S"
	HIGH_CARD_SAME_HIGHEST_LOWER_THIRD  string = "3S 4H 4S 8S 9S"
	HIGH_CARD_SAME_HIGHEST_LOWER_FOURTH string = "3S 3H 5S 8S 9S"
	HIGH_CARD_LOWEST                    string = "2S 2H 2D 2C 3S"
	P1_WINS_INPUT                       string = "3S 4H 5S 8S 9S 2S 2H 2D 2C 3S"
)


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
