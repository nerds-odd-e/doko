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
	sortRank := []string{"A", "K", "Q"}

	for _, v := range sortRank {
		hasQueen := strings.Index(card, v)
		if hasQueen < 14 && hasQueen >= 0 {
			return true
		}
	}

	hasQueen := strings.Index(card, "Q")
	if hasQueen < 14 && hasQueen >= 0 {
		return true
	}
	hasAce := strings.Index(card, "A")
	if hasAce < 14 && hasAce >= 0 {
		return true
	}
	hasKing := strings.Index(card, "K")
	if hasKing < 14 && hasKing >= 0 {
		return true
	}
	return false
}
