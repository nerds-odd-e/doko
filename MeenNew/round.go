package meennew

import "strings"

type Round struct {
	value string
}

func (r *Round) isPlayer1Winner() bool {
	sortRank := []string{"A", "K", "Q"}
	for _, rank := range sortRank {
		founded := strings.Index(r.value, rank)
		if foundAtPlayer1(founded) {
			return true
		}
		if founded >= 14 && founded >= -1 {
			return false
		}
	}
	return false
}

func foundAtPlayer1(founded int) bool {
	return founded < 14 && founded >= 0
}
