package poker

import (
	"strings"
)

func IsOnePair(cards []string) bool {
	cardRanks := []string{"A", "K"}

	for _, cardRank := range cardRanks {
		counter := 0

		for _, card := range cards {
			counter += strings.Count(card, cardRank)
		}

		if counter == 2 {
			return true
		}
	}

	return false

}
