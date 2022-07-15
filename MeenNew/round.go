package main

import (
	"strings"
)

type Round struct {
	value string
}

func (r Round) isPlayer1Winner() bool {
	sortRank := []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}
	for _, rank := range sortRank {
		if strings.Count(r.value[0:14], rank) == 2 {
			return true
		}
		if strings.Count(r.value, rank) > 1 {
			return false
		}
		if strings.Index(r.value, rank) < 0 {
			continue
		}
		return strings.Index(r.value, rank) < 14
	}
	return false
}
