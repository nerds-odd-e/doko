package meennew

import "strings"

type Round struct {
	value string
}

func (r *Round) isPlayer1Winner() bool {
	sortRank := []string{"A", "K", "Q"}
	for _, rank := range sortRank {
		founded := strings.Index(r.value, rank)
		if founded < 0 {
			continue
		}
		if foundAtPlayer1(founded) {
			return true
		}
		if founded >= 14 {
			return false
		}
	}
	return false
}

func foundAtPlayer1(founded int) bool {
	return founded < 14 && founded >= 0
}
