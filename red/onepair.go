package poker

import (
	"strings"
)

func IsOnePair(cards []string) bool {
	counter := 0

	for _, card := range cards {
		counter += strings.Count(card, "A")
	}

	return counter == 2
}
