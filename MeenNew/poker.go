package meennew

import "strings"

func PokerHand(records []string) int {

	countP1Winner := 0
	if len(records) == 0 {
		return countP1Winner
	}
	for _, v := range records {
		if isPlayer1Winner(v) {
			countP1Winner += 1
		}
	}
	return countP1Winner
}

func isPlayer1Winner(round string) bool {
	sortRank := []string{"A", "K", "Q"}
	for _, rank := range sortRank {
		hasQueen := strings.Index(round, rank)
		if hasQueen < 14 && hasQueen >= 0 {
			return true
		}
		if hasQueen >= 14 {
			return false
		}
	}
	return false
}
