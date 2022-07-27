package poker

import (
	"strings"
)

func IsOnePair(cards []string) bool {
	counterA := 0

	for _, card := range cards {
		counterA += strings.Count(card, "A")
	}

	if counterA == 2 {
		return true
	}

	counterK := 0

	for _, card := range cards {
		counterK += strings.Count(card, "K")
	}

	return counterK == 2

}
