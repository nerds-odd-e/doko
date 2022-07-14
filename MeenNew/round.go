package meennew

import "strings"

type Round struct {
	value string
}

func (r *Round) isPlayer1Winner() bool {
	sortRank := []string{"A", "K", "Q"}
	for _, rank := range sortRank {
		founded := strings.Index(r.value, rank)
		foundAtPlayer1 := founded < 14 && founded >= 0
		if foundAtPlayer1 {
			return true
		}
		if founded >= 14 {
			return false
		}
	}
	return false
}
