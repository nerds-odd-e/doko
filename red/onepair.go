package poker

import (
	"strings"
)

func IsOnePair(cards []string) bool {
	// rules := [][]string {
	// 	{"AC", "KS", "2D", "3H", "5S"},
	// 	{"TC", "AS", "2D", "3H", "5S"},
	// }
	// for _, rule := range rules {
	// 	if reflect.DeepEqual(cards, rule) {
	// 		return false
	// 	}
	// }

	counter := 0

	for _, card := range cards {
		counter += strings.Count(card, "A")
	}

	return counter == 2
}
