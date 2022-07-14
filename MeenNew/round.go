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
		return foundAtPlayer1(founded)
	}
	return false
}

func foundAtPlayer1(founded int) bool {
	return founded < 14
}
