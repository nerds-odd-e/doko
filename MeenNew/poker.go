package meennew

import "strings"

func PokerHand(records []string) int {

	countP1Winner := 0
	if len(records) == 0 {
		return countP1Winner
	}
	for _, v := range records {
		if isWinner(v) {
			countP1Winner += 1
		}
	}
	return countP1Winner
}

func isWinner(card string) bool {
	hasAce := strings.Index(card, "A")
	if hasAce < 14 && hasAce >= 0 {
		return true
	}
	hasTen := strings.Index(card, "K")
	if hasTen < 14 && hasTen >= 0 {
		return true
	}
	return false
}
