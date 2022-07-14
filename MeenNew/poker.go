package meennew

import "strings"

type Round struct {
	value string
}

func PokerHand(records []string) int {

	countP1Winner := 0
	if len(records) == 0 {
		return countP1Winner
	}
	for _, value := range records {
		round := Round{value}
		if round.isPlayer1Winner() {
			countP1Winner += 1
		}
	}
	return countP1Winner
}

func (r *Round) isPlayer1Winner() bool {
	sortRank := []string{"A", "K", "Q"}
	for _, rank := range sortRank {
		founded := strings.Index(r.value, rank)
		if founded < 14 && founded >= 0 {
			return true
		}
		if founded >= 14 {
			return false
		}
	}
	return false
}
